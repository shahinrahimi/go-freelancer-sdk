package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type ProjectsProjectsService struct {
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

func (s *ProjectsProjectsService) Do(ctx context.Context) (*ResponseProjects, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "projects/0.1/projects/",
	}
	if s.projects != nil {
		r.addParam("projects", s.projects)
	}
	if s.owners != nil {
		r.addParam("owners", s.owners)
	}
	if s.bidders != nil {
		r.addParam("bidders", s.bidders)
	}
	if s.seoUrls != nil {
		r.addParam("seo_urls", s.seoUrls)
	}
	if s.fromTime != 0 {
		r.addParam("from_time", s.fromTime)
	}
	if s.toTime != 0 {
		r.addParam("to_time", s.toTime)
	}
	if s.frontendProjectStatuses != nil {
		r.addParam("frontend_project_statuses", s.frontendProjectStatuses)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := new(ResponseProjects)
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (s *ProjectsProjectsService) SetProjects(projects []int64) *ProjectsProjectsService {
	s.projects = projects
	return s
}

func (s *ProjectsProjectsService) SetOwners(owners []int64) *ProjectsProjectsService {
	s.owners = owners
	return s
}

func (s *ProjectsProjectsService) SetBidders(bidders []int64) *ProjectsProjectsService {
	s.bidders = bidders
	return s
}

func (s *ProjectsProjectsService) SetSeoUrls(seoUrls []string) *ProjectsProjectsService {
	s.seoUrls = seoUrls
	return s
}

func (s *ProjectsProjectsService) SetFromTime(fromTime int64) *ProjectsProjectsService {
	s.fromTime = fromTime
	return s
}

func (s *ProjectsProjectsService) SetToTime(toTime int64) *ProjectsProjectsService {
	s.toTime = toTime
	return s
}

func (s *ProjectsProjectsService) SetFullDescription(fullDescription bool) *ProjectsProjectsService {
	s.fullDescription = fullDescription
	return s
}

func (s *ProjectsProjectsService) SetJobDetails(jobDetails bool) *ProjectsProjectsService {
	s.jobDetails = jobDetails
	return s
}

func (s *ProjectsProjectsService) SetUpgradeDetails(upgradeDetails bool) *ProjectsProjectsService {
	s.upgradeDetails = upgradeDetails
	return s
}

func (s *ProjectsProjectsService) SetAttachmentDetails(attachmentDetails bool) *ProjectsProjectsService {
	s.attachmentDetails = attachmentDetails
	return s
}

func (s *ProjectsProjectsService) SetFileDetails(fileDetails bool) *ProjectsProjectsService {
	s.fileDetails = fileDetails
	return s
}

func (s *ProjectsProjectsService) SetQualificationDetails(qualificationDetails bool) *ProjectsProjectsService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *ProjectsProjectsService) SetSelectedBids(selectedBids bool) *ProjectsProjectsService {
	s.selectedBids = selectedBids
	return s
}

func (s *ProjectsProjectsService) SetHireMeDetails(hireMeDetails bool) *ProjectsProjectsService {
	s.hireMeDetails = hireMeDetails
	return s
}

func (s *ProjectsProjectsService) SetUserDetails(userDetails bool) *ProjectsProjectsService {
	s.userDetails = userDetails
	return s
}

func (s *ProjectsProjectsService) SetInvitedFreelancerDetails(invitedFreelancerDetails bool) *ProjectsProjectsService {
	s.invitedFreelancerDetails = invitedFreelancerDetails
	return s
}

func (s *ProjectsProjectsService) SetRecommendedFreelancerDetails(recommendedFreelancerDetails bool) *ProjectsProjectsService {
	s.recommendedFreelancerDetails = recommendedFreelancerDetails
	return s
}

func (s *ProjectsProjectsService) SetSupportSessionDetails(supportSessionDetails bool) *ProjectsProjectsService {
	s.supportSessionDetails = supportSessionDetails
	return s
}

func (s *ProjectsProjectsService) SetLocationDetails(locationDetails bool) *ProjectsProjectsService {
	s.locationDetails = locationDetails
	return s
}

func (s *ProjectsProjectsService) SetNdaSignatureDetails(ndaSignatureDetails bool) *ProjectsProjectsService {
	s.ndaSignatureDetails = ndaSignatureDetails
	return s
}

func (s *ProjectsProjectsService) SetDriveFileDetails(driveFileDetails bool) *ProjectsProjectsService {
	s.driveFileDetails = driveFileDetails
	return s
}

func (s *ProjectsProjectsService) SetNdaDetails(ndaDetails bool) *ProjectsProjectsService {
	s.ndaDetails = ndaDetails
	return s
}

func (s *ProjectsProjectsService) SetLocalDetails(localDetails bool) *ProjectsProjectsService {
	s.localDetails = localDetails
	return s
}

func (s *ProjectsProjectsService) SetEquipmentDetails(equipmentDetails bool) *ProjectsProjectsService {
	s.equipmentDetails = equipmentDetails
	return s
}

func (s *ProjectsProjectsService) SetClientEngagementDetails(clientEngagementDetails bool) *ProjectsProjectsService {
	s.clientEngagementDetails = clientEngagementDetails
	return s
}

func (s *ProjectsProjectsService) SetUserResponsiveness(userResponsiveness bool) *ProjectsProjectsService {
	s.userResponsiveness = userResponsiveness
	return s
}

func (s *ProjectsProjectsService) SetServiceOfferingDetails(serviceOfferingDetails bool) *ProjectsProjectsService {
	s.serviceOfferingDetails = serviceOfferingDetails
	return s
}

func (s *ProjectsProjectsService) SetCorporateUsers(corporateUsers bool) *ProjectsProjectsService {
	s.corporateUsers = corporateUsers
	return s
}

func (s *ProjectsProjectsService) SetIsNonHireMe(isNonHireMe bool) *ProjectsProjectsService {
	s.isNonHireMe = isNonHireMe
	return s
}

func (s *ProjectsProjectsService) SetHasMilestone(hasMilestone bool) *ProjectsProjectsService {
	s.hasMilestone = hasMilestone
	return s
}

func (s *ProjectsProjectsService) SetTeam(team bool) *ProjectsProjectsService {
	s.team = team
	return s
}

func (s *ProjectsProjectsService) SetCompact(compact bool) *ProjectsProjectsService {
	s.compact = compact
	return s
}

func (s *ProjectsProjectsService) SetLimit(limit int) *ProjectsProjectsService {
	s.limit = limit
	return s
}

func (s *ProjectsProjectsService) SetOffset(offset int) *ProjectsProjectsService {
	s.offset = offset
	return s
}

func (s *ProjectsProjectsService) SetFrontendProjectStatuses(frontendProjectStatuses []string) *ProjectsProjectsService {
	s.frontendProjectStatuses = frontendProjectStatuses
	return s
}
func (s *ProjectsProjectsService) SetProximityDetails(proximityDetails bool) *ProjectsProjectsService {
	s.proximityDetails = proximityDetails
	return s
}

func (s *ProjectsProjectsService) SetReviewAvailabilityDetails(reviewAvailabilityDetails bool) *ProjectsProjectsService {
	s.reviewAvailabilityDetails = reviewAvailabilityDetails
	return s
}
func (s *ProjectsProjectsService) SetNegotiatedDetails(negotiatedDetails bool) *ProjectsProjectsService {
	s.negotiatedDetails = negotiatedDetails
	return s
}
func (s *ProjectsProjectsService) SetUserAvatar(userAvatar bool) *ProjectsProjectsService {
	s.userAvatar = userAvatar
	return s
}
func (s *ProjectsProjectsService) SetUserCountryDetails(userCountryDetails bool) *ProjectsProjectsService {
	s.userCountryDetails = userCountryDetails
	return s
}
func (s *ProjectsProjectsService) SetUserProfileDescription(userProfileDescription bool) *ProjectsProjectsService {
	s.userProfileDescription = userProfileDescription
	return s
}

func (s *ProjectsProjectsService) SetProjectCollaborationDetails(projectCollaborationDetails bool) *ProjectsProjectsService {
	s.projectCollaborationDetails = projectCollaborationDetails
	return s
}

func (s *ProjectsProjectsService) SetUserDisplayInfo(userDisplayInfo bool) *ProjectsProjectsService {
	s.userDisplayInfo = userDisplayInfo
	return s
}

func (s *ProjectsProjectsService) SetUserJobs(userJobs bool) *ProjectsProjectsService {
	s.userJobs = userJobs
	return s
}

func (s *ProjectsProjectsService) SetUserBalanceDetails(userBalanceDetails bool) *ProjectsProjectsService {
	s.userBalanceDetails = userBalanceDetails
	return s
}
func (s *ProjectsProjectsService) SetUserQualificationDetails(userQualificationDetails bool) *ProjectsProjectsService {
	s.userQualificationDetails = userQualificationDetails
	return s
}

func (s *ProjectsProjectsService) SetUserMembershipDetails(userMembershipDetails bool) *ProjectsProjectsService {
	s.userMembershipDetails = userMembershipDetails
	return s
}

func (s *ProjectsProjectsService) SetUserFinancialDetails(userFinancialDetails bool) *ProjectsProjectsService {
	s.userFinancialDetails = userFinancialDetails
	return s
}

func (s *ProjectsProjectsService) SetUserLocationDetails(userLocationDetails bool) *ProjectsProjectsService {
	s.userLocationDetails = userLocationDetails
	return s
}

func (s *ProjectsProjectsService) SetUserPortfolioDetails(userPortfolioDetails bool) *ProjectsProjectsService {
	s.userPortfolioDetails = userPortfolioDetails
	return s
}

func (s *ProjectsProjectsService) SetUserPreferredDetails(userPreferredDetails bool) *ProjectsProjectsService {
	s.userPreferredDetails = userPreferredDetails
	return s
}

func (s *ProjectsProjectsService) SetUserBadgeDetails(userBadgeDetails bool) *ProjectsProjectsService {
	s.userBadgeDetails = userBadgeDetails
	return s
}

func (s *ProjectsProjectsService) SetUserStatus(userStatus bool) *ProjectsProjectsService {
	s.userStatus = userStatus
	return s
}

func (s *ProjectsProjectsService) SetUserReputation(userReputation bool) *ProjectsProjectsService {
	s.userReputation = userReputation
	return s
}

func (s *ProjectsProjectsService) SetUserEmployerReputation(userEmployerReputation bool) *ProjectsProjectsService {
	s.userEmployerReputation = userEmployerReputation
	return s
}

func (s *ProjectsProjectsService) SetUserEmployerReputationExtra(userEmployerReputationExtra bool) *ProjectsProjectsService {
	s.userEmployerReputationExtra = userEmployerReputationExtra
	return s
}

func (s *ProjectsProjectsService) SetUserCoverImage(userCoverImage bool) *ProjectsProjectsService {
	s.userCoverImage = userCoverImage
	return s
}

func (s *ProjectsProjectsService) SetUserPastCoverImage(userPastCoverImage bool) *ProjectsProjectsService {
	s.userPastCoverImage = userPastCoverImage
	return s
}

func (s *ProjectsProjectsService) SetUserRecommendations(userRecommendations bool) *ProjectsProjectsService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *ProjectsProjectsService) SetMarketingMobileNumber(marketingMobileNumber bool) *ProjectsProjectsService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *ProjectsProjectsService) SetSanctionDetails(sanctionDetails bool) *ProjectsProjectsService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *ProjectsProjectsService) SetLimitedAccount(limitedAccount bool) *ProjectsProjectsService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *ProjectsProjectsService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *ProjectsProjectsService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}
