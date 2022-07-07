resource "genesyscloud_journey_action_map" "terraform_test_-TEST-CASE-" {
  display_name          = "terraform_test_-TEST-CASE-_updated"
  trigger_with_segments = [genesyscloud_journey_segment.terraform_test_-TEST-CASE-_action_map_dependency.id]
  activation {
    type = "immediate"
  }
  action {
    media_type = "architectFlow"
    architect_flow_fields {
      architect_flow_id = "1e5fe2dc-9973-42b7-a328-c015617b3a98" # This is a random hardcoded value!
      flow_request_mappings {
        name           = "Name_1"
        attribute_type = "String"
        mapping_type   = "Lookup"
        value          = "session.id"
      }
      flow_request_mappings {
        name           = "Name_2"
        attribute_type = "Integer"
        mapping_type   = "HardCoded"
        value          = "999"
      }
    }
  }
  start_date = "2022-07-04T12:00:00.000000"

  depends_on = [genesyscloud_journey_segment.terraform_test_-TEST-CASE-_action_map_dependency]
}

resource "genesyscloud_journey_segment" "terraform_test_-TEST-CASE-_action_map_dependency" {
  display_name            = "terraform_test_-TEST-CASE-_action_map_dependency"
  color                   = "#008000"
  scope                   = "Customer"
  should_display_to_agent = false
  external_segment {
    id     = "4654654654"
    name   = "external segment name"
    source = "AdobeExperiencePlatform"
  }
}
