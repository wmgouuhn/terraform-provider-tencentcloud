package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func init() {
	// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=tencentcloud_ci_media_speech_recognition_template
	resource.AddTestSweepers("tencentcloud_ci_media_speech_recognition_template", &resource.Sweeper{
		Name: "tencentcloud_ci_media_speech_recognition_template",
		F: func(r string) error {
			logId := getLogId(contextNil)
			ctx := context.WithValue(context.TODO(), logIdKey, logId)
			cli, _ := sharedClientForRegion(r)
			client := cli.(*TencentCloudClient).apiV3Conn
			service := CiService{client: client}

			response, _, err := service.client.UseCiClient(defaultCiBucket).CI.DescribeMediaTemplate(ctx, &cos.DescribeMediaTemplateOptions{
				Name: "speech_recognition_template",
			})
			if err != nil {
				return err
			}

			for _, v := range response.TemplateList {
				err := service.DeleteCiMediaTemplateById(ctx, defaultCiBucket, v.TemplateId)
				if err != nil {
					continue
				}
			}

			return nil
		},
	})
}

// go test -i; go test -test.run TestAccTencentCloudCiMediaSpeechRecognitionTemplateResource_basic -v
func TestAccTencentCloudCiMediaSpeechRecognitionTemplateResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiMediaSpeechRecognitionTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiMediaSpeechRecognitionTemplate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiMediaSpeechRecognitionTemplateExists("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template"),
					resource.TestCheckResourceAttrSet("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "id"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "bucket", defaultCiBucket),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "name", "speech_recognition_template"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.#", "1"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.engine_model_type", "16k_zh"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.channel_num", "1"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.res_text_format", "1"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.filter_dirty", "0"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.filter_modal", "1"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.convert_num_mode", "0"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.speaker_diarization", "1"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.speaker_number", "0"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.filter_punc", "0"),
					resource.TestCheckResourceAttr("tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template", "speech_recognition.0.output_file_type", "txt"),
				),
			},
			{
				ResourceName:      "tencentcloud_ci_media_speech_recognition_template.media_speech_recognition_template",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCiMediaSpeechRecognitionTemplateDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := CiService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloud_ci_media_speech_recognition_template" {
			continue
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		bucket := idSplit[0]
		templateId := idSplit[1]

		res, err := service.DescribeCiMediaTemplateById(ctx, bucket, templateId)
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("ci media speech recognition template still exist, Id: %v\n", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckCiMediaSpeechRecognitionTemplateExists(re string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		service := CiService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		rs, ok := s.RootModule().Resources[re]
		if !ok {
			return fmt.Errorf("ci media speech recognition template %s is not found", re)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf(" id is not set")
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		bucket := idSplit[0]
		templateId := idSplit[1]

		result, err := service.DescribeCiMediaTemplateById(ctx, bucket, templateId)
		if err != nil {
			return err
		}

		if result == nil {
			return fmt.Errorf("ci media speech recognition template not found, Id: %v", rs.Primary.ID)
		}

		return nil
	}
}

const testAccCiMediaSpeechRecognitionTemplateVar = `
variable "bucket" {
	default = "` + defaultCiBucket + `"
  }

`

const testAccCiMediaSpeechRecognitionTemplate = testAccCiMediaSpeechRecognitionTemplateVar + `

resource "tencentcloud_ci_media_speech_recognition_template" "media_speech_recognition_template" {
	bucket = var.bucket
	name = "speech_recognition_template"
	speech_recognition {
		engine_model_type = "16k_zh"
		channel_num = "1"
		res_text_format = "1"
		filter_dirty = "0"
		filter_modal = "1"
		convert_num_mode = "0"
		speaker_diarization = "1"
		speaker_number = "0"
		filter_punc = "0"
		output_file_type = "txt"
	}
}

`