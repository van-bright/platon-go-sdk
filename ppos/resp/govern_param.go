package resp

type ParamItem struct {
	/**
	 * 参数模块
	 */
	Module string `json:"Module"`
	/**
	 * 参数名称
	 */
	Name string `json:"Name"`
	/**
	 * 参数说明
	 */
	Desc string `json:"Desc"`
}

type ParamValue struct {

	/**
	 * 旧参数值
	 */
	StaleValue string `json:"StaleValue"`
	/**
	 * 参数值
	 */
	Value string `json:"Value"`
	/**
	 * 块高。(>=ActiveBLock，将取Value;否则取StaleValue)
	 */
	ActiveBlock uint64 `json:"ActiveBlock"`
}

type GovernParam struct {
	ParamItem  ParamItem  `json:"ParamItem"`
	ParamValue ParamValue `json:"ParamValue"`
}
