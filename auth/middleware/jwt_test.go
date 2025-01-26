package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJWTMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config := JWTConfig{
		SecretKey:     "test-secret",
		TokenDuration: time.Hour,
	}

	t.Run("Valid Token", func(t *testing.T) {
		token, err := GenerateToken("123", "testuser", []string{"user"}, config)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET("/test", JWTMiddleware(config), func(c *gin.Context) {
			claims, exists := c.Get("claims")
			assert.True(t, exists)
			assert.NotNil(t, claims)
			c.Status(http.StatusOK)
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Missing Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET("/test", JWTMiddleware(config), func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Invalid Token Format", func(t *testing.T) {
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET("/test", JWTMiddleware(config), func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "InvalidFormat")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Expired Token", func(t *testing.T) {
		config := JWTConfig{
			SecretKey:     "test-secret",
			TokenDuration: -time.Hour, // Expired
		}

		token, err := GenerateToken("123", "testuser", []string{"user"}, config)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		r.GET("/test", JWTMiddleware(config), func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestGenerateToken(t *testing.T) {
	config := JWTConfig{
		SecretKey:     "test-secret",
		TokenDuration: time.Hour,
	}

	token, err := GenerateToken("123", "testuser", []string{"user"}, config)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
