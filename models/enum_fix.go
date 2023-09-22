package models

func (x AffiliationType) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *AffiliationType) UnmarshalText(text []byte) error { *x = AffiliationType(text); return nil }

func (x ConsentPurpose) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *ConsentPurpose) UnmarshalText(text []byte) error { *x = ConsentPurpose(text); return nil }

func (x ErrorCode) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *ErrorCode) UnmarshalText(text []byte) error { *x = ErrorCode(text); return nil }

func (x Gender) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *Gender) UnmarshalText(text []byte) error { *x = Gender(text); return nil }

func (x GrantType) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *GrantType) UnmarshalText(text []byte) error { *x = GrantType(text); return nil }

func (x GroupType) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *GroupType) UnmarshalText(text []byte) error { *x = GroupType(text); return nil }

func (x MaritalStatus) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *MaritalStatus) UnmarshalText(text []byte) error { *x = MaritalStatus(text); return nil }

func (x OrgType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}
func (x *OrgType) UnmarshalText(text []byte) error {
	*x = OrgType(text)
	return nil
}

func (x OrgActiveStatus) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}
func (x *OrgActiveStatus) UnmarshalText(text []byte) error {
	*x = OrgActiveStatus(text)
	return nil
}

func (x PersonRelationType) MarshalText() ([]byte, error) { return []byte(string(x)), nil }
func (x *PersonRelationType) UnmarshalText(text []byte) error {
	*x = PersonRelationType(text)
	return nil
}

func (x RelationType) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *RelationType) UnmarshalText(text []byte) error { *x = RelationType(text); return nil }

func (x RoleScope) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *RoleScope) UnmarshalText(text []byte) error { *x = RoleScope(text); return nil }

func (x SearchVisibility) MarshalText() ([]byte, error)     { return []byte(string(x)), nil }
func (x *SearchVisibility) UnmarshalText(text []byte) error { *x = SearchVisibility(text); return nil }
