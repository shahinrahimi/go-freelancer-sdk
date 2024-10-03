package freelancer

import (
	"context"
	"net/http"
)

type ProjectService struct {
	client                       *Client
	projects                     []int64
	owners                       []int64
	bidders                      []int64
	seoUrls                      []string
	fromTime                     int64
	toTime                       int64
	frontendProjectStatuses      []string
	team                         bool
	isNonHireMe                  bool
	hasMilestone                 bool
	fullDescription              bool
	jobDetails                   bool
	upgradeDetails               bool
	attachmentDetails            bool
	fileDetails                  bool
	qualificationDetails         bool
	selectedBids                 bool
	hireMeDetails                bool
	userDetails                  bool
	invitedFreelancerDetails     bool
	recommendedFreelancerDetails bool
	supportSessionDetails        bool
	locationDetails              bool
	ndaSignatureDetails          bool
	projectCollaborationDetails  bool
	proximityDetails             bool
	reviewAvailabilityDetails    bool
	negotiatedDetails            bool
	driveFileDetails             bool
	ndaDetails                   bool
	localDetails                 bool
	equipmentDetails             bool
	clientEngagementDetails      bool
	serviceOfferingDetails       bool
	userAvatar                   bool
	userCountryDetails           bool
	userProfileDescription       bool
	userDisplayInfo              bool
	userJobs                     bool
	userBalanceDetails           bool
	userQualificationDetails     bool
	userMembershipDetails        bool
	userFinancialDetails         bool
	userLocationDetails          bool
	userPortfolioDetails         bool
	userPreferredDetails         bool
	userBadgeDetails             bool
	userStatus                   bool
	userReputation               bool
	userEmployerReputation       bool
	userEmployerReputationExtra  bool
	userCoverImage               bool
	userPastCoverImage           bool
	userRecommendations          bool
	userResponsiveness           bool
	corporateUsers               bool
	marketingMobileNumber        bool
	sanctionDetails              bool
	limitedAccount               bool
	equipmentGroupDetails        bool
	limit                        int
	offset                       int
	compact                      bool
}

func (s *ProjectService) Do(ctx context.Context) (*http.Response, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/projects/0.1/projects/",
	}
	m := params{}
	if s.attachmentDetails != nil {

	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (s *ProjectService) GetProjects(id int) (*Project, *http.Response, error) {
	return nil, nil, nil
}

func (s *ProjectService) Projects(projects []int64) *ProjectService {
	s.projects = projects
	return s
}

func (s *ProjectService) Owners(owners []int64) *ProjectService {
	s.owners = owners
	return s
}

func (s *ProjectService) Bidders(bidders []int64) *ProjectService {
	s.bidders = bidders
	return s
}

func (s *ProjectService) SeoUrls(seoUrls []string) *ProjectService {
	s.seoUrls = seoUrls
	return s
}

func (s *ProjectService) FromTime(fromTime int64) *ProjectService {
	s.fromTime = fromTime
	return s
}

func (s *ProjectService) ToTime(toTime int64) *ProjectService {
	s.toTime = toTime
	return s
}

func (s *ProjectService) FullDescription(fullDescription bool) *ProjectService {
	s.fullDescription = fullDescription
	return s
}

func (s *ProjectService) JobDetails(jobDetails bool) *ProjectService {
	s.jobDetails = jobDetails
	return s
}

func (s *ProjectService) UpgradeDetails(upgradeDetails bool) *ProjectService {
	s.upgradeDetails = upgradeDetails
	return s
}

func (s *ProjectService) AttachmentDetails(attachmentDetails bool) *ProjectService {
	s.attachmentDetails = attachmentDetails
	return s
}

func (s *ProjectService) FileDetails(fileDetails bool) *ProjectService {
	s.fileDetails = fileDetails
	return s
}

func (s *ProjectService) QualificationDetails(qualificationDetails bool) *ProjectService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *ProjectService) SelectedBids(selectedBids bool) *ProjectService {
	s.selectedBids = selectedBids
	return s
}

func (s *ProjectService) HireMeDetails(hireMeDetails bool) *ProjectService {
	s.hireMeDetails = hireMeDetails
	return s
}

func (s *ProjectService) UserDetails(userDetails bool) *ProjectService {
	s.userDetails = userDetails
	return s
}

func (s *ProjectService) InvitedFreelancerDetails(invitedFreelancerDetails bool) *ProjectService {
	s.invitedFreelancerDetails = invitedFreelancerDetails
	return s
}

func (s *ProjectService) RecommendedFreelancerDetails(recommendedFreelancerDetails bool) *ProjectService {
	s.recommendedFreelancerDetails = recommendedFreelancerDetails
	return s
}

func (s *ProjectService) SupportSessionDetails(supportSessionDetails bool) *ProjectService {
	s.supportSessionDetails = supportSessionDetails
	return s
}

func (s *ProjectService) LocationDetails(locationDetails bool) *ProjectService {
	s.locationDetails = locationDetails
	return s
}

func (s *ProjectService) NdaSignatureDetails(ndaSignatureDetails bool) *ProjectService {
	s.ndaSignatureDetails = ndaSignatureDetails
	return s
}

func (s *ProjectService) DriveFileDetails(driveFileDetails bool) *ProjectService {
	s.driveFileDetails = driveFileDetails
	return s
}

func (s *ProjectService) NdaDetails(ndaDetails bool) *ProjectService {
	s.ndaDetails = ndaDetails
	return s
}

func (s *ProjectService) LocalDetails(localDetails bool) *ProjectService {
	s.localDetails = localDetails
	return s
}

func (s *ProjectService) EquipmentDetails(equipmentDetails bool) *ProjectService {
	s.equipmentDetails = equipmentDetails
	return s
}

func (s *ProjectService) ClientEngagementDetails(clientEngagementDetails bool) *ProjectService {
	s.clientEngagementDetails = clientEngagementDetails
	return s
}

func (s *ProjectService) UserResponsiveness(userResponsiveness bool) *ProjectService {
	s.userResponsiveness = userResponsiveness
	return s
}

func (s *ProjectService) ServiceOfferingDetails(serviceOfferingDetails bool) *ProjectService {
	s.serviceOfferingDetails = serviceOfferingDetails
	return s
}

func (s *ProjectService) CorporateUsers(corporateUsers bool) *ProjectService {
	s.corporateUsers = corporateUsers
	return s
}

func (s *ProjectService) IsNonHireMe(isNonHireMe bool) *ProjectService {
	s.isNonHireMe = isNonHireMe
	return s
}

func (s *ProjectService) HasMilestone(hasMilestone bool) *ProjectService {
	s.hasMilestone = hasMilestone
	return s
}

func (s *ProjectService) Team(team bool) *ProjectService {
	s.team = team
	return s
}

func (s *ProjectService) Compact(compact bool) *ProjectService {
	s.compact = compact
	return s
}

func (s *ProjectService) Limit(limit int) *ProjectService {
	s.limit = limit
	return s
}

func (s *ProjectService) Offset(offset int) *ProjectService {
	s.offset = offset
	return s
}

func (s *ProjectService) FrontendProjectStatuses(frontendProjectStatuses []string) *ProjectService {
	s.frontendProjectStatuses = frontendProjectStatuses
	return s
}
func (s *ProjectService) ProximityDetails(proximityDetails bool) *ProjectService {
	s.proximityDetails = proximityDetails
	return s
}

func (s *ProjectService) ReviewAvailabilityDetails(reviewAvailabilityDetails bool) *ProjectService {
	s.reviewAvailabilityDetails = reviewAvailabilityDetails
	return s
}
func (s *ProjectService) NegotiatedDetails(negotiatedDetails bool) *ProjectService {
	s.negotiatedDetails = negotiatedDetails
	return s
}
func (s *ProjectService) UserAvatar(userAvatar bool) *ProjectService {
	s.userAvatar = userAvatar
	return s
}
func (s *ProjectService) UserCountryDetails(userCountryDetails bool) *ProjectService {
	s.userCountryDetails = userCountryDetails
	return s
}
func (s *ProjectService) UserProfileDescription(userProfileDescription bool) *ProjectService {
	s.userProfileDescription = userProfileDescription
	return s
}

func (s *ProjectService) ProjectCollaborationDetails(projectCollaborationDetails bool) *ProjectService {
	s.projectCollaborationDetails = projectCollaborationDetails
	return s
}

func (s *ProjectService) UserDisplayInfo(userDisplayInfo bool) *ProjectService {
	s.userDisplayInfo = userDisplayInfo
	return s
}

func (s *ProjectService) UserJobs(userJobs bool) *ProjectService {
	s.userJobs = userJobs
	return s
}

func (s *ProjectService) UserBalanceDetails(userBalanceDetails bool) *ProjectService {
	s.userBalanceDetails = userBalanceDetails
	return s
}
func (s *ProjectService) UserQualificationDetails(userQualificationDetails bool) *ProjectService {
	s.userQualificationDetails = userQualificationDetails
	return s
}

func (s *ProjectService) UserMembershipDetails(userMembershipDetails bool) *ProjectService {
	s.userMembershipDetails = userMembershipDetails
	return s
}

func (s *ProjectService) UserFinancialDetails(userFinancialDetails bool) *ProjectService {
	s.userFinancialDetails = userFinancialDetails
	return s
}

func (s *ProjectService) UserLocationDetails(userLocationDetails bool) *ProjectService {
	s.userLocationDetails = userLocationDetails
	return s
}

func (s *ProjectService) UserPortfolioDetails(userPortfolioDetails bool) *ProjectService {
	s.userPortfolioDetails = userPortfolioDetails
	return s
}

func (s *ProjectService) UserPreferredDetails(userPreferredDetails bool) *ProjectService {
	s.userPreferredDetails = userPreferredDetails
	return s
}

func (s *ProjectService) UserBadgeDetails(userBadgeDetails bool) *ProjectService {
	s.userBadgeDetails = userBadgeDetails
	return s
}

func (s *ProjectService) UserStatus(userStatus bool) *ProjectService {
	s.userStatus = userStatus
	return s
}

func (s *ProjectService) UserReputation(userReputation bool) *ProjectService {
	s.userReputation = userReputation
	return s
}

func (s *ProjectService) UserEmployerReputation(userEmployerReputation bool) *ProjectService {
	s.userEmployerReputation = userEmployerReputation
	return s
}

func (s *ProjectService) UserEmployerReputationExtra(userEmployerReputationExtra bool) *ProjectService {
	s.userEmployerReputationExtra = userEmployerReputationExtra
	return s
}

func (s *ProjectService) UserCoverImage(userCoverImage bool) *ProjectService {
	s.userCoverImage = userCoverImage
	return s
}

func (s *ProjectService) UserPastCoverImage(userPastCoverImage bool) *ProjectService {
	s.userPastCoverImage = userPastCoverImage
	return s
}

func (s *ProjectService) UserRecommendations(userRecommendations bool) *ProjectService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ProjectService) MarketingMobileNumber(marketingMobileNumber bool) *ProjectService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ProjectService) SanctionDetails(sanctionDetails bool) *ProjectService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ProjectService) LimitedAccount(limitedAccount bool) *ProjectService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ProjectService) EquipmentGroupDetails(equipmentGroupDetails bool) *ProjectService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}
