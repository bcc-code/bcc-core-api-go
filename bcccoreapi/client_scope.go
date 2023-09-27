package bcccoreapi

type Scope string

const (
	ScopePersonsRead               Scope = "persons#read"
	ScopePersonsNameRead           Scope = "persons.name#read"
	ScopePersonsBirthDateRead      Scope = "persons.birth_date#read"
	ScopePersonsDeceasedDateRead   Scope = "persons.deceased_date#read"
	ScopePersonsNationalIdsRead    Scope = "persons.national_ids#read"
	ScopePersonsGenderRead         Scope = "persons.gender#read"
	ScopePersonsEmailRead          Scope = "persons.email#read"
	ScopePersonsPhoneRead          Scope = "persons.phone#read"
	ScopePersonsPreferencesRead    Scope = "persons.preferences#read"
	ScopePersonsAddressRead        Scope = "persons.address#read"
	ScopePersonsProfilePictureRead Scope = "persons.profile_picture#read"
	ScopePersonAffiliationsRead    Scope = "persons.affiliations#read"
	ScopePersonRelationsRead       Scope = "persons.relations#read"
	ScopePersonConsentsRead        Scope = "persons.consents#read"
	ScopePersonRoleAssignmentsRead Scope = "persons.role_assignments#read"
	ScopePersonPersonIdRead        Scope = "persons.person_id#read"
	ScopeOrgsRead                  Scope = "orgs#read"
	ScopeCountriesRead             Scope = "countries#read"
	ScopeRolesRead                 Scope = "roles#read"
	ScopeGroupsRead                Scope = "groups#read"

	ScopePersonsWrite               Scope = "persons#write"
	ScopePersonAffiliationsWrite    Scope = "persons.affiliations#write"
	ScopePersonRelationsWrite       Scope = "persons.relations#write"
	ScopePersonConsentsWrite        Scope = "persons.consents#write"
	ScopePersonRoleAssignmentsWrite Scope = "persons.role_assignments#write"
	ScopeOrgsWrite                  Scope = "orgs#write"
	ScopeCountriesWrite             Scope = "countries#write"
	ScopeRolesWrite                 Scope = "roles#write"
	ScopeGroupsWrite                Scope = "groups#write"
)
