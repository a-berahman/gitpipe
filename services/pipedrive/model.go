package pipedrive

//Error implemented from pipedrive error schema
type Error struct {
	Status    bool   `json:"status"`
	Error     string `json:"error"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
}

// AddActivityRQ implemented from pipedrive Add an Activity request schema
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/addActivity
type AddActivityRQ struct {
	DueDate string `json:"due_date"`
	Note    string `json:"note"`
	Subject string `json:"subject"`
}

// AddActivityRS implemented from pipedrive Add an Activity response schema
// https://developers.pipedrive.com/docs/api/v1/#!/Activities/addActivity
type AddActivityRS struct {
	Success bool `json:"success"`
	Data    struct {
		ID                         int         `json:"id"`
		CompanyID                  int         `json:"company_id"`
		UserID                     int         `json:"user_id"`
		Done                       bool        `json:"done"`
		Type                       string      `json:"type"`
		ReferenceType              interface{} `json:"reference_type"`
		ReferenceID                interface{} `json:"reference_id"`
		ConferenceMeetingClient    interface{} `json:"conference_meeting_client"`
		ConferenceMeetingURL       interface{} `json:"conference_meeting_url"`
		DueDate                    string      `json:"due_date"`
		DueTime                    string      `json:"due_time"`
		Duration                   string      `json:"duration"`
		BusyFlag                   interface{} `json:"busy_flag"`
		AddTime                    string      `json:"add_time"`
		MarkedAsDoneTime           string      `json:"marked_as_done_time"`
		LastNotificationTime       interface{} `json:"last_notification_time"`
		LastNotificationUserID     interface{} `json:"last_notification_user_id"`
		NotificationLanguageID     interface{} `json:"notification_language_id"`
		Subject                    string      `json:"subject"`
		PublicDescription          interface{} `json:"public_description"`
		CalendarSyncIncludeContext interface{} `json:"calendar_sync_include_context"`
		Location                   interface{} `json:"location"`
		OrgID                      interface{} `json:"org_id"`
		PersonID                   interface{} `json:"person_id"`
		DealID                     interface{} `json:"deal_id"`
		LeadID                     interface{} `json:"lead_id"`
		LeadTitle                  string      `json:"lead_title"`
		ActiveFlag                 bool        `json:"active_flag"`
		UpdateTime                 string      `json:"update_time"`
		UpdateUserID               interface{} `json:"update_user_id"`
		GcalEventID                interface{} `json:"gcal_event_id"`
		GoogleCalendarID           interface{} `json:"google_calendar_id"`
		GoogleCalendarEtag         interface{} `json:"google_calendar_etag"`
		SourceTimezone             interface{} `json:"source_timezone"`
		RecRule                    interface{} `json:"rec_rule"`
		RecRuleExtension           interface{} `json:"rec_rule_extension"`
		RecMasterActivityID        interface{} `json:"rec_master_activity_id"`
		ConferenceMeetingID        interface{} `json:"conference_meeting_id"`
		Note                       string      `json:"note"`
		CreatedByUserID            int         `json:"created_by_user_id"`
		LocationSubpremise         interface{} `json:"location_subpremise"`
		LocationStreetNumber       interface{} `json:"location_street_number"`
		LocationRoute              interface{} `json:"location_route"`
		LocationSublocality        interface{} `json:"location_sublocality"`
		LocationLocality           interface{} `json:"location_locality"`
		LocationAdminAreaLevel1    interface{} `json:"location_admin_area_level_1"`
		LocationAdminAreaLevel2    interface{} `json:"location_admin_area_level_2"`
		LocationCountry            interface{} `json:"location_country"`
		LocationPostalCode         interface{} `json:"location_postal_code"`
		LocationFormattedAddress   interface{} `json:"location_formatted_address"`
		Attendees                  interface{} `json:"attendees"`
		Participants               interface{} `json:"participants"`
		Series                     interface{} `json:"series"`
		OrgName                    interface{} `json:"org_name"`
		PersonName                 interface{} `json:"person_name"`
		DealTitle                  interface{} `json:"deal_title"`
		OwnerName                  string      `json:"owner_name"`
		PersonDropboxBcc           interface{} `json:"person_dropbox_bcc"`
		DealDropboxBcc             interface{} `json:"deal_dropbox_bcc"`
		AssignedToUserID           int         `json:"assigned_to_user_id"`
		TypeName                   string      `json:"type_name"`
		File                       interface{} `json:"file"`
	} `json:"data"`
	AdditionalData struct {
		UpdatesStoryID interface{} `json:"updates_story_id"`
	} `json:"additional_data"`
	RelatedObjects struct {
		User struct {
			Num12023814 struct {
				ID         int         `json:"id"`
				Name       string      `json:"name"`
				Email      string      `json:"email"`
				HasPic     int         `json:"has_pic"`
				PicHash    interface{} `json:"pic_hash"`
				ActiveFlag bool        `json:"active_flag"`
			} `json:"12023814"`
		} `json:"user"`
	} `json:"related_objects"`
	Error Error
}
