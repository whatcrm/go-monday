package monday

//func (c *Get) Contacts(contactID string) (out []models.ContactResult, err error) {
//	options := callMethodOptions{
//		Method:  fiber.MethodPost,
//		BaseURL: CrmContactGet,
//		In: &RequestParams{
//			ID: contactID,
//		},
//		Out:    &models.Contact{},
//		Params: nil,
//	}
//
//	if contactID != "" {
//		if err = c.api.callMethod(options); err != nil {
//			return
//		}
//		out = []models.ContactResult{options.Out.(*models.Contact).Result}
//	}
//
//	if contactID == "" {
//		options.In = nil
//		options.BaseURL = CrmContactList
//		options.Out = &models.ContactList{}
//		if err = c.api.callMethod(options); err != nil {
//			return
//		}
//		out = options.Out.(*models.ContactList).Result
//	}
//	return
//}
