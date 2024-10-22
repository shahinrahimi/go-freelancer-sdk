package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectActiveService struct {
	client                      *Client
	query                       string
	projectTypes                []ProjectType
	projectUpgrades             []ProjectUpgradeType
	sortFields                  []SortFieldsType
	contestUpgrades             []ContestUpgradeType
	minAvgPrice                 float64
	maxAvgPrice                 float64
	minAvgHourlyRate            float64
	maxAvgHourlyRate            float64
	minPrice                    float64
	maxPrice                    float64
	minHourlyRate               float64
	maxHourlyRate               float64
	jobs                        []int64
	countries                   []string
	languages                   []string
	latitude                    float64
	longitude                   float64
	fromTime                    int64
	toTime                      int64
	sortField                   SortFieldsType
	projectIDs                  []int64
	topRightLatitude            float64
	topRightLongitude           float64
	bottomLeftLatitude          float64
	bottomLeftLongitude         float64
	reverseSort                 bool
	orSearchQuery               string
	highlightPreTags            string
	highlightPostTags           string
	fullDescription             bool
	jobDetails                  bool
	upgradeDetails              bool
	userDetails                 bool
	locationDetails             bool
	ndaSignatureDetails         bool
	projectCollaborationDetails bool
	userAvatar                  bool
	userCountryDetails          bool
	userProfileDescription      bool
	userDisplayInfo             bool
	userJobs                    bool
	userBalanceDetails          bool
	userQualificationDetails    bool
	userMembershipDetails       bool
	userFinancialDetails        bool
	userLocationDetails         bool
	userPortfolioDetails        bool
	userPreferredDetails        bool
	userBadgeDetails            bool
	userStatus                  bool
	userReputation              bool
	userEmployerReputation      bool
	userReputationExtra         bool
	userEmployerReputationExtra bool
	userCoverImage              bool
	userPastCoverImage          bool
	userRecommendations         bool
	userResponsiveness          bool
	corporateUsers              bool
	marketingMobileNumber       bool
	sanctionDetails             bool
	limitedAccount              bool
	equipmentGroupDetails       bool
	limit                       int
	offset                      int
	compact                     bool
}

func (c *Client) NewProjectActiveService() *ProjectActiveService {
	return &ProjectActiveService{client: c}
}

func (s *ProjectActiveService) Do(ctx context.Context) (*ResponseProjects, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "projects/0.1/projects/active/",
	}
	if s.query != "" {
		r.addParam("query", s.query)
	}
	if s.projectTypes != nil {
		r.addParam("project_types", s.projectTypes)
	}
	if s.projectUpgrades != nil {
		r.addParam("project_upgrades", s.projectUpgrades)
	}
	if s.sortFields != nil {
		r.addParam("sort_fields", s.sortFields)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseProjects{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProjectActiveService) SetUserResponsiveness(userResponsiveness bool) *ProjectActiveService {
	s.userResponsiveness = userResponsiveness
	return s
}

func (s *ProjectActiveService) SetCorporateUsers(corporateUsers bool) *ProjectActiveService {
	s.corporateUsers = corporateUsers
	return s
}

func (s *ProjectActiveService) SetMarketingMobileNumber(marketingMobileNumber bool) *ProjectActiveService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ProjectActiveService) SetSanctionDetails(sanctionDetails bool) *ProjectActiveService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ProjectActiveService) SetLimitedAccount(limitedAccount bool) *ProjectActiveService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ProjectActiveService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ProjectActiveService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *ProjectActiveService) SetUserDetails(userDetails bool) *ProjectActiveService {
	s.userDetails = userDetails
	return s
}
func (s *ProjectActiveService) SetLocationDetails(locationDetails bool) *ProjectActiveService {
	s.locationDetails = locationDetails
	return s
}

func (s *ProjectActiveService) SetNdaSignatureDetails(ndaSignatureDetails bool) *ProjectActiveService {
	s.ndaSignatureDetails = ndaSignatureDetails
	return s
}

func (s *ProjectActiveService) SetProjectCollaborationDetails(projectCollaborationDetails bool) *ProjectActiveService {
	s.projectCollaborationDetails = projectCollaborationDetails
	return s
}

func (s *ProjectActiveService) SetUserAvatar(userAvatar bool) *ProjectActiveService {
	s.userAvatar = userAvatar
	return s
}

func (s *ProjectActiveService) SetUserCountryDetails(userCountryDetails bool) *ProjectActiveService {
	s.userCountryDetails = userCountryDetails
	return s
}

func (s *ProjectActiveService) SetUserProfileDescription(userProfileDescription bool) *ProjectActiveService {
	s.userProfileDescription = userProfileDescription
	return s
}

func (s *ProjectActiveService) SetUserDisplayInfo(userDisplayInfo bool) *ProjectActiveService {
	s.userDisplayInfo = userDisplayInfo
	return s
}

func (s *ProjectActiveService) SetUserJobs(userJobs bool) *ProjectActiveService {
	s.userJobs = userJobs
	return s
}

func (s *ProjectActiveService) SetUserBalanceDetails(userBalanceDetails bool) *ProjectActiveService {
	s.userBalanceDetails = userBalanceDetails
	return s
}

func (s *ProjectActiveService) SetUserQualificationDetails(userQualificationDetails bool) *ProjectActiveService {
	s.userQualificationDetails = userQualificationDetails
	return s
}

func (s *ProjectActiveService) SetUserMembershipDetails(userMembershipDetails bool) *ProjectActiveService {
	s.userMembershipDetails = userMembershipDetails
	return s
}

func (s *ProjectActiveService) SetUserFinancialDetails(userFinancialDetails bool) *ProjectActiveService {
	s.userFinancialDetails = userFinancialDetails
	return s
}

func (s *ProjectActiveService) SetUserLocationDetails(userLocationDetails bool) *ProjectActiveService {
	s.userLocationDetails = userLocationDetails
	return s
}

func (s *ProjectActiveService) SetUserPortfolioDetails(userPortfolioDetails bool) *ProjectActiveService {
	s.userPortfolioDetails = userPortfolioDetails
	return s
}

func (s *ProjectActiveService) SetUserPreferredDetails(userPreferredDetails bool) *ProjectActiveService {
	s.userPreferredDetails = userPreferredDetails
	return s
}

func (s *ProjectActiveService) SetUserBadgeDetails(userBadgeDetails bool) *ProjectActiveService {
	s.userBadgeDetails = userBadgeDetails
	return s
}

func (s *ProjectActiveService) SetUserStatus(userStatus bool) *ProjectActiveService {
	s.userStatus = userStatus
	return s
}

func (s *ProjectActiveService) SetUserReputation(userReputation bool) *ProjectActiveService {
	s.userReputation = userReputation
	return s
}

func (s *ProjectActiveService) SetUserEmployerReputation(userEmployerReputation bool) *ProjectActiveService {
	s.userEmployerReputation = userEmployerReputation
	return s
}

func (s *ProjectActiveService) SetUserReputationExtra(userReputationExtra bool) *ProjectActiveService {
	s.userReputationExtra = userReputationExtra
	return s
}

func (s *ProjectActiveService) SetUserEmployerReputationExtra(userEmployerReputationExtra bool) *ProjectActiveService {
	s.userEmployerReputationExtra = userEmployerReputationExtra
	return s
}

func (s *ProjectActiveService) SetUserCoverImage(userCoverImage bool) *ProjectActiveService {
	s.userCoverImage = userCoverImage
	return s
}

func (s *ProjectActiveService) SetUserPastCoverImage(userPastCoverImage bool) *ProjectActiveService {
	s.userPastCoverImage = userPastCoverImage
	return s
}

func (s *ProjectActiveService) SetUserRecommendations(userRecommendations bool) *ProjectActiveService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ProjectActiveService) SetProjectTypes(projectTypes []ProjectType) *ProjectActiveService {
	s.projectTypes = projectTypes
	return s
}

func (s *ProjectActiveService) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *ProjectActiveService {
	s.projectUpgrades = projectUpgrades
	return s
}

func (s *ProjectActiveService) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *ProjectActiveService {
	s.contestUpgrades = contestUpgrades
	return s
}
func (s *ProjectActiveService) SetCompact(compact bool) *ProjectActiveService {
	s.compact = compact
	return s
}
func (s *ProjectActiveService) SetLimit(limit int) *ProjectActiveService {
	s.limit = limit
	return s
}

func (s *ProjectActiveService) SetOffset(offset int) *ProjectActiveService {
	s.offset = offset
	return s
}

func (s *ProjectActiveService) SetReverseSort(reverseSort bool) *ProjectActiveService {
	s.reverseSort = reverseSort
	return s
}

func (s *ProjectActiveService) SetFullDescription(fullDescription bool) *ProjectActiveService {
	s.fullDescription = fullDescription
	return s
}

func (s *ProjectActiveService) SetJobDetails(jobDetails bool) *ProjectActiveService {
	s.jobDetails = jobDetails
	return s
}

func (s *ProjectActiveService) SetUpgradeDetails(upgradeDetails bool) *ProjectActiveService {
	s.upgradeDetails = upgradeDetails
	return s
}

func (s *ProjectActiveService) SetQuery(query string) *ProjectActiveService {
	s.query = query
	return s
}

func (s *ProjectActiveService) SetMinAvgPrice(minAvgPrice float64) *ProjectActiveService {
	s.minAvgPrice = minAvgPrice
	return s
}

func (s *ProjectActiveService) SetMaxAvgPrice(maxAvgPrice float64) *ProjectActiveService {
	s.maxAvgPrice = maxAvgPrice
	return s
}

func (s *ProjectActiveService) SetMinAvgHourlyRate(minAvgHourlyRate float64) *ProjectActiveService {
	s.minAvgHourlyRate = minAvgHourlyRate
	return s
}

func (s *ProjectActiveService) SetMaxAvgHourlyRate(maxAvgHourlyRate float64) *ProjectActiveService {
	s.maxAvgHourlyRate = maxAvgHourlyRate
	return s
}

func (s *ProjectActiveService) SetMinPrice(minPrice float64) *ProjectActiveService {
	s.minPrice = minPrice
	return s
}

func (s *ProjectActiveService) SetMaxPrice(maxPrice float64) *ProjectActiveService {
	s.maxPrice = maxPrice
	return s
}

func (s *ProjectActiveService) SetMinHourlyRate(minHourlyRate float64) *ProjectActiveService {
	s.minHourlyRate = minHourlyRate
	return s
}

func (s *ProjectActiveService) SetMaxHourlyRate(maxHourlyRate float64) *ProjectActiveService {
	s.maxHourlyRate = maxHourlyRate
	return s
}

func (s *ProjectActiveService) SetJobs(jobs []int64) *ProjectActiveService {
	s.jobs = jobs
	return s
}

func (s *ProjectActiveService) SetCountries(countries []string) *ProjectActiveService {
	s.countries = countries
	return s
}

func (s *ProjectActiveService) SetLanguages(languages []string) *ProjectActiveService {
	s.languages = languages
	return s
}

func (s *ProjectActiveService) SetLatitude(latitude float64) *ProjectActiveService {
	s.latitude = latitude
	return s
}

func (s *ProjectActiveService) SetLongitude(longitude float64) *ProjectActiveService {
	s.longitude = longitude
	return s
}

func (s *ProjectActiveService) SetFromTime(fromTime int64) *ProjectActiveService {
	s.fromTime = fromTime
	return s
}

func (s *ProjectActiveService) SetToTime(toTime int64) *ProjectActiveService {
	s.toTime = toTime
	return s
}

func (s *ProjectActiveService) SetSortField(sortField SortFieldsType) *ProjectActiveService {
	s.sortField = sortField
	return s
}

func (s *ProjectActiveService) SetProjectIDs(projectIDs []int64) *ProjectActiveService {
	s.projectIDs = projectIDs
	return s
}

func (s *ProjectActiveService) SetTopRightLatitude(topRightLatitude float64) *ProjectActiveService {
	s.topRightLatitude = topRightLatitude
	return s
}

func (s *ProjectActiveService) SetTopRightLongitude(topRightLongitude float64) *ProjectActiveService {
	s.topRightLongitude = topRightLongitude
	return s
}

func (s *ProjectActiveService) SetBottomLeftLatitude(bottomLeftLatitude float64) *ProjectActiveService {
	s.bottomLeftLatitude = bottomLeftLatitude
	return s
}

func (s *ProjectActiveService) SetBottomLeftLongitude(bottomLeftLongitude float64) *ProjectActiveService {
	s.bottomLeftLongitude = bottomLeftLongitude
	return s
}

func (s *ProjectActiveService) SetOrSearchQuery(orSearchQuery string) *ProjectActiveService {
	s.orSearchQuery = orSearchQuery
	return s
}

func (s *ProjectActiveService) SetHighlightPreTags(highlightPreTags string) *ProjectActiveService {
	s.highlightPreTags = highlightPreTags
	return s
}

func (s *ProjectActiveService) SetHighlightPostTags(highlightPostTags string) *ProjectActiveService {
	s.highlightPostTags = highlightPostTags
	return s
}
