package go_project_template_test

import (
	"testing"

	go_project_template "github.com/chyroc/go-project-template"
	"github.com/stretchr/testify/assert"
)

func Test_Incr(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		as.Equal(2, go_project_template.Incr(1))
	})
}
