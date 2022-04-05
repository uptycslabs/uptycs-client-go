## Uptycs Client (Go)

A Go library for [uptycs](https://uptycs.io)

### Examples ###

#### AlertRules ####

Get all alert rules

```
	rules, _ := c.GetAlertRules()
	for _, alert := range rules.Items {
	  if alert.Name == "Wide Open SG Test" {
	    fmt.Println(fmt.Sprintf("%s has id %s", alert.Name, alert.ID))
	  }
	}
```

Get an alert rule by ID

```
	rule, _ := c.GetAlertRule(uptycs.AlertRule{
		ID: "1d4720ce-19a9-4a03-bb3a-905717b8a60f",
	})
	fmt.Println(fmt.Sprintf("Found rule with name %s as %s", rule.Name, rule.ID))
```

Create an alert rule

```
	rule, err := c.CreateAlertRule(uptycs.AlertRule{
	  Name: "marcus test",
	  Description: "marcus test",
	  Grouping: "MITRE",
	  GroupingL2: "Impact",
	  GroupingL3: "T1560",
	  SQLConfig: uptycs.SQLConfig{
	    IntervalSeconds: 3600,
	  },
	  Code: "test_marc",
	  Type: "sql",
	  Rule: "select * from processes limit 1 :to;",
	})
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(fmt.Sprintf("Created Rule '%s' with id %s", rule.Name, rule.ID))
```

Update an alert rule by by ID

```
	updatedRule, err := c.UpdateAlertRule(uptycs.AlertRule{
	  ID: "2b23acf7-6c3f-4ff9-8c98-039551a9300d",
	  Name: "marcus test updated",
	  Description: "marcus test updated",
	  Grouping: "MITRE",
	  GroupingL2: "Impact",
	  GroupingL3: "T1580",
	  SQLConfig: uptycs.SQLConfig{
	    IntervalSeconds: 1800,
	  },
	  Code: "test_marc2",
	  Type: "sql",
	  Rule: "select * from processes limit 2 :to;",
	})
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println(fmt.Sprintf("Updated Rule '%s' with id %s", updatedRule.Name, updatedRule.ID))
```

Delete an alert rule by ID

```
	rule := uptycs.AlertRule{
	  Name: "marcus test updated",
	  ID: "2b23acf7-6c3f-4ff9-8c98-039551a9300d",
	}

	_, err = c.DeleteAlertRule(uptycs.AlertRule{
	  ID: rule.ID,
	})
	fmt.Println(fmt.Sprintf("Deleted Rule '%s' with id %s", rule.Name, rule.ID))
```

