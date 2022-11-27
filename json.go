package drmaa2interface

import "encoding/json"

func (jt *JobTemplate) UnmarshalJSON(data []byte) error {
	type jtType JobTemplate
	var jt2 jtType
	if err := json.Unmarshal(data, &jt2); err != nil {
		return err
	}
	*jt = JobTemplate(jt2)
	return nil
}

func (jt JobTemplate) MarshalJSON() ([]byte, error) {
	type jtType JobTemplate
	return json.Marshal(jtType(jt))
}
