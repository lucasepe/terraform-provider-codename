package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceCodename_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCodename_basic,
				Check: resource.ComposeTestCheckFunc(
					dumpCodename(t, "codename.id_1"),
					testAccResourceCodenameParts("codename.id_1", "-", 2),
				),
			},
		},
	})
}

func TestAccResourceCodename_token_length(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCodename_snakefy,
				Check: resource.ComposeTestCheckFunc(
					dumpCodename(t, "codename.id_1"),
					testAccResourceCodenameParts("codename.id_1", "_", 3),
					testAccResourceCodenameTokenLength("codename.id_1", "_", 4),
				),
			},
		},
	})
}

func TestAccResourceCodename_prefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCodename_prefix,
				Check: resource.ComposeTestCheckFunc(
					dumpCodename(t, "codename.id_1"),
					testAccResourceCodenameParts("codename.id_1", "-", 2),
					resource.TestMatchResourceAttr(
						"codename.id_1", "id", regexp.MustCompile("^it_")),
				),
			},
		},
	})
}

func testAccResourceCodenameParts(id string, separator string, atLeast int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found: %s", id)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		parts := strings.Split(rs.Primary.ID, separator)
		if got := len(parts); got < atLeast {
			return fmt.Errorf("codename fewer parts than expected; got [%d] want >= [%d]", got, atLeast)
		}

		return nil
	}
}

func testAccResourceCodenameTokenLength(id string, separator string, tokenLength int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found: %s", id)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		parts := strings.Split(rs.Primary.ID, separator)
		if got := len(parts[len(parts)-1]); got != tokenLength {
			return fmt.Errorf("token length does not match got [%d], want [%d]", got, tokenLength)
		}

		return nil
	}
}

func dumpCodename(t *testing.T, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[id]
		if !ok {
			return fmt.Errorf("Not found: %s", id)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		t.Logf(rs.Primary.ID)

		return nil
	}
}

const testAccResourceCodename_basic = `
resource "codename" "id_1" {
}
`

const testAccResourceCodename_snakefy = `
resource "codename" "id_1" {
  snakefy = true
  token_length = 4
}
`
const testAccResourceCodename_prefix = `
resource "codename" "id_1" {
  prefix = "it_"
}
`
