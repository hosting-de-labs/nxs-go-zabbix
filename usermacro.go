package zabbix

import ()

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermacro/object#global_macro */
type GlobalmacroObject struct {
	GlobalmacroID string `json:"globalmacroid,omitempty"`
	Macro         string `json:"macro ,omitempty"`
	Value         string `json:"value ,omitempty"`
}

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermacro/object#host_macro */
type HostmacroObject struct {
	HostmacroID string `json:"hostmacroid,omitempty"`
	HostID      string `json:"hostid,omitempty"`
	Macro       string `json:"macro,omitempty"`
	Value       string `json:"value,omitempty"`

	Groups    []HostgroupObject `json:"groups,omitempty"`
	Hosts     []HostObject      `json:"hosts,omitempty"`
	Templates []TemplateObject  `json:"templates,omitempty"`
}

/* see: https://www.zabbix.com/documentation/2.4/manual/api/reference/usermacro/get#parameters */
type HostmacroGetParams struct {
	GetParameters

	GroupIDs     []string `json:"groupids,omitempty"`
	HostIDs      []string `json:"hostids,omitempty"`
	HostmacroIDs []string `json:"hostmacroids,omitempty"`
	TemplateIDs  []string `json:"templateids,omitempty"`

	SelectGroups    SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

/* Structure to store creation result */
type hostmacroCreateResult struct {
	HostmacroIDs []string `json:"hostmacroids"`
}

/* Structure to store deletion result */
type hostmacroDeleteResult struct {
	HostmacroIDs []string `json:"hostmacroids"`
}

func (z *Zabbix) HostmacroGet(params HostmacroGetParams) ([]HostmacroObject, int, error) {

	var result []HostmacroObject

	status, err := z.request("usermacro.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Zabbix) HostmacroCreate(params []HostmacroObject) ([]string, int, error) {

	var result hostmacroCreateResult

	status, err := z.request("usermacro.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostmacroIDs, status, nil
}

func (z *Zabbix) HostmacroDelete(hostmacroIDs []string) ([]string, int, error) {

	var result hostmacroDeleteResult

	status, err := z.request("usermacro.delete", hostmacroIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostmacroIDs, status, nil
}
