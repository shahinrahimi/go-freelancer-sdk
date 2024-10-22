package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectsActiveService struct {
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

func (s *ProjectsActiveService) Do(ctx context.Context) (*ResponseProjects, error) {
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

func (s *ProjectsActiveService) SetUserResponsiveness(userResponsiveness bool) *ProjectsActiveService {
	s.userResponsiveness = userResponsiveness
	return s
}

func (s *ProjectsActiveService) SetCorporateUsers(corporateUsers bool) *ProjectsActiveService {
	s.corporateUsers = corporateUsers
	return s
}

func (s *ProjectsActiveService) SetMarketingMobileNumber(marketingMobileNumber bool) *ProjectsActiveService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ProjectsActiveService) SetSanctionDetails(sanctionDetails bool) *ProjectsActiveService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ProjectsActiveService) SetLimitedAccount(limitedAccount bool) *ProjectsActiveService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ProjectsActiveService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ProjectsActiveService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *ProjectsActiveService) SetUserDetails(userDetails bool) *ProjectsActiveService {
	s.userDetails = userDetails
	return s
}
func (s *ProjectsActiveService) SetLocationDetails(locationDetails bool) *ProjectsActiveService {
	s.locationDetails = locationDetails
	return s
}

func (s *ProjectsActiveService) SetNdaSignatureDetails(ndaSignatureDetails bool) *ProjectsActiveService {
	s.ndaSignatureDetails = ndaSignatureDetails
	return s
}

func (s *ProjectsActiveService) SetProjectCollaborationDetails(projectCollaborationDetails bool) *ProjectsActiveService {
	s.projectCollaborationDetails = projectCollaborationDetails
	return s
}

func (s *ProjectsActiveService) SetUserAvatar(userAvatar bool) *ProjectsActiveService {
	s.userAvatar = userAvatar
	return s
}

func (s *ProjectsActiveService) SetUserCountryDetails(userCountryDetails bool) *ProjectsActiveService {
	s.userCountryDetails = userCountryDetails
	return s
}

func (s *ProjectsActiveService) SetUserProfileDescription(userProfileDescription bool) *ProjectsActiveService {
	s.userProfileDescription = userProfileDescription
	return s
}

func (s *ProjectsActiveService) SetUserDisplayInfo(userDisplayInfo bool) *ProjectsActiveService {
	s.userDisplayInfo = userDisplayInfo
	return s
}

func (s *ProjectsActiveService) SetUserJobs(userJobs bool) *ProjectsActiveService {
	s.userJobs = userJobs
	return s
}

func (s *ProjectsActiveService) SetUserBalanceDetails(userBalanceDetails bool) *ProjectsActiveService {
	s.userBalanceDetails = userBalanceDetails
	return s
}

func (s *ProjectsActiveService) SetUserQualificationDetails(userQualificationDetails bool) *ProjectsActiveService {
	s.userQualificationDetails = userQualificationDetails
	return s
}

func (s *ProjectsActiveService) SetUserMembershipDetails(userMembershipDetails bool) *ProjectsActiveService {
	s.userMembershipDetails = userMembershipDetails
	return s
}

func (s *ProjectsActiveService) SetUserFinancialDetails(userFinancialDetails bool) *ProjectsActiveService {
	s.userFinancialDetails = userFinancialDetails
	return s
}

func (s *ProjectsActiveService) SetUserLocationDetails(userLocationDetails bool) *ProjectsActiveService {
	s.userLocationDetails = userLocationDetails
	return s
}

func (s *ProjectsActiveService) SetUserPortfolioDetails(userPortfolioDetails bool) *ProjectsActiveService {
	s.userPortfolioDetails = userPortfolioDetails
	return s
}

func (s *ProjectsActiveService) SetUserPreferredDetails(userPreferredDetails bool) *ProjectsActiveService {
	s.userPreferredDetails = userPreferredDetails
	return s
}

func (s *ProjectsActiveService) SetUserBadgeDetails(userBadgeDetails bool) *ProjectsActiveService {
	s.userBadgeDetails = userBadgeDetails
	return s
}

func (s *ProjectsActiveService) SetUserStatus(userStatus bool) *ProjectsActiveService {
	s.userStatus = userStatus
	return s
}

func (s *ProjectsActiveService) SetUserReputation(userReputation bool) *ProjectsActiveService {
	s.userReputation = userReputation
	return s
}

func (s *ProjectsActiveService) SetUserEmployerReputation(userEmployerReputation bool) *ProjectsActiveService {
	s.userEmployerReputation = userEmployerReputation
	return s
}

func (s *ProjectsActiveService) SetUserReputationExtra(userReputationExtra bool) *ProjectsActiveService {
	s.userReputationExtra = userReputationExtra
	return s
}

func (s *ProjectsActiveService) SetUserEmployerReputationExtra(userEmployerReputationExtra bool) *ProjectsActiveService {
	s.userEmployerReputationExtra = userEmployerReputationExtra
	return s
}

func (s *ProjectsActiveService) SetUserCoverImage(userCoverImage bool) *ProjectsActiveService {
	s.userCoverImage = userCoverImage
	return s
}

func (s *ProjectsActiveService) SetUserPastCoverImage(userPastCoverImage bool) *ProjectsActiveService {
	s.userPastCoverImage = userPastCoverImage
	return s
}

func (s *ProjectsActiveService) SetUserRecommendations(userRecommendations bool) *ProjectsActiveService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ProjectsActiveService) SetProjectTypes(projectTypes []ProjectType) *ProjectsActiveService {
	s.projectTypes = projectTypes
	return s
}

func (s *ProjectsActiveService) SetProjectUpgrades(projectUpgrades []ProjectUpgradeType) *ProjectsActiveService {
	s.projectUpgrades = projectUpgrades
	return s
}

func (s *ProjectsActiveService) SetContestUpgrades(contestUpgrades []ContestUpgradeType) *ProjectsActiveService {
	s.contestUpgrades = contestUpgrades
	return s
}
func (s *ProjectsActiveService) SetCompact(compact bool) *ProjectsActiveService {
	s.compact = compact
	return s
}
func (s *ProjectsActiveService) SetLimit(limit int) *ProjectsActiveService {
	s.limit = limit
	return s
}

func (s *ProjectsActiveService) SetOffset(offset int) *ProjectsActiveService {
	s.offset = offset
	return s
}

func (s *ProjectsActiveService) SetReverseSort(reverseSort bool) *ProjectsActiveService {
	s.reverseSort = reverseSort
	return s
}

func (s *ProjectsActiveService) SetFullDescription(fullDescription bool) *ProjectsActiveService {
	s.fullDescription = fullDescription
	return s
}

func (s *ProjectsActiveService) SetJobDetails(jobDetails bool) *ProjectsActiveService {
	s.jobDetails = jobDetails
	return s
}

func (s *ProjectsActiveService) SetUpgradeDetails(upgradeDetails bool) *ProjectsActiveService {
	s.upgradeDetails = upgradeDetails
	return s
}

func (s *ProjectsActiveService) SetQuery(query string) *ProjectsActiveService {
	s.query = query
	return s
}

func (s *ProjectsActiveService) SetMinAvgPrice(minAvgPrice float64) *ProjectsActiveService {
	s.minAvgPrice = minAvgPrice
	return s
}

func (s *ProjectsActiveService) SetMaxAvgPrice(maxAvgPrice float64) *ProjectsActiveService {
	s.maxAvgPrice = maxAvgPrice
	return s
}

func (s *ProjectsActiveService) SetMinAvgHourlyRate(minAvgHourlyRate float64) *ProjectsActiveService {
	s.minAvgHourlyRate = minAvgHourlyRate
	return s
}

func (s *ProjectsActiveService) SetMaxAvgHourlyRate(maxAvgHourlyRate float64) *ProjectsActiveService {
	s.maxAvgHourlyRate = maxAvgHourlyRate
	return s
}

func (s *ProjectsActiveService) SetMinPrice(minPrice float64) *ProjectsActiveService {
	s.minPrice = minPrice
	return s
}

func (s *ProjectsActiveService) SetMaxPrice(maxPrice float64) *ProjectsActiveService {
	s.maxPrice = maxPrice
	return s
}

func (s *ProjectsActiveService) SetMinHourlyRate(minHourlyRate float64) *ProjectsActiveService {
	s.minHourlyRate = minHourlyRate
	return s
}

func (s *ProjectsActiveService) SetMaxHourlyRate(maxHourlyRate float64) *ProjectsActiveService {
	s.maxHourlyRate = maxHourlyRate
	return s
}

func (s *ProjectsActiveService) SetJobs(jobs []int64) *ProjectsActiveService {
	s.jobs = jobs
	return s
}

func (s *ProjectsActiveService) SetCountries(countries []string) *ProjectsActiveService {
	s.countries = countries
	return s
}

func (s *ProjectsActiveService) SetLanguages(languages []string) *ProjectsActiveService {
	s.languages = languages
	return s
}

func (s *ProjectsActiveService) SetLatitude(latitude float64) *ProjectsActiveService {
	s.latitude = latitude
	return s
}

func (s *ProjectsActiveService) SetLongitude(longitude float64) *ProjectsActiveService {
	s.longitude = longitude
	return s
}

func (s *ProjectsActiveService) SetFromTime(fromTime int64) *ProjectsActiveService {
	s.fromTime = fromTime
	return s
}

func (s *ProjectsActiveService) SetToTime(toTime int64) *ProjectsActiveService {
	s.toTime = toTime
	return s
}

func (s *ProjectsActiveService) SetSortField(sortField SortFieldsType) *ProjectsActiveService {
	s.sortField = sortField
	return s
}

func (s *ProjectsActiveService) SetProjectIDs(projectIDs []int64) *ProjectsActiveService {
	s.projectIDs = projectIDs
	return s
}

func (s *ProjectsActiveService) SetTopRightLatitude(topRightLatitude float64) *ProjectsActiveService {
	s.topRightLatitude = topRightLatitude
	return s
}

func (s *ProjectsActiveService) SetTopRightLongitude(topRightLongitude float64) *ProjectsActiveService {
	s.topRightLongitude = topRightLongitude
	return s
}

func (s *ProjectsActiveService) SetBottomLeftLatitude(bottomLeftLatitude float64) *ProjectsActiveService {
	s.bottomLeftLatitude = bottomLeftLatitude
	return s
}

func (s *ProjectsActiveService) SetBottomLeftLongitude(bottomLeftLongitude float64) *ProjectsActiveService {
	s.bottomLeftLongitude = bottomLeftLongitude
	return s
}

func (s *ProjectsActiveService) SetOrSearchQuery(orSearchQuery string) *ProjectsActiveService {
	s.orSearchQuery = orSearchQuery
	return s
}

func (s *ProjectsActiveService) SetHighlightPreTags(highlightPreTags string) *ProjectsActiveService {
	s.highlightPreTags = highlightPreTags
	return s
}

func (s *ProjectsActiveService) SetHighlightPostTags(highlightPostTags string) *ProjectsActiveService {
	s.highlightPostTags = highlightPostTags
	return s
}
