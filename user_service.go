package freelancer

type UserService struct {
	client                        *Client
	users                         []int64
	usernames                     []string
	avatar                        bool
	countryDetails                bool
	profileDescription            bool
	displayInfo                   bool
	jobs                          bool
	balanceDetails                bool
	qualificationDetails          bool
	membershipDetails             bool
	financialDetails              bool
	locationDetails               bool
	portfolioDetails              bool
	preferredDetails              bool
	badgeDetails                  bool
	status                        bool
	reputation                    bool
	employerReputation            bool
	reputationExtra               bool
	employerReputationExtra       bool
	coverImage                    bool
	pastCoverImage                bool
	mobileTracking                bool
	bidQualityDetails             bool
	depositMethods                bool
	userRecommendations           bool
	marketingMobileNumber         bool
	sanctionDetails               bool
	limitedAccount                bool
	completedUserRelevantJobCount bool
	equipmentGroupDetails         bool
	jobRanks                      bool
	jobSeoDetails                 bool
	risingStar                    bool
	shareholderDetails            bool
	staffDetails                  bool
}

func (s *UserService) Users(users []int64) *UserService {
	s.users = users
	return s
}

func (s *UserService) Usernames(usernames []string) *UserService {
	s.usernames = usernames
	return s
}

func (s *UserService) Avatar(avatar bool) *UserService {
	s.avatar = avatar
	return s
}

func (s *UserService) CountryDetails(countryDetails bool) *UserService {
	s.countryDetails = countryDetails
	return s
}

func (s *UserService) ProfileDescription(profileDescription bool) *UserService {
	s.profileDescription = profileDescription
	return s
}

func (s *UserService) DisplayInfo(displayInfo bool) *UserService {
	s.displayInfo = displayInfo
	return s
}

func (s *UserService) Jobs(jobs bool) *UserService {
	s.jobs = jobs
	return s
}

func (s *UserService) BalanceDetails(balanceDetails bool) *UserService {
	s.balanceDetails = balanceDetails
	return s
}

func (s *UserService) QualificationDetails(qualificationDetails bool) *UserService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *UserService) MembershipDetails(membershipDetails bool) *UserService {
	s.membershipDetails = membershipDetails
	return s
}

func (s *UserService) FinancialDetails(financialDetails bool) *UserService {
	s.financialDetails = financialDetails
	return s
}

func (s *UserService) LocationDetails(locationDetails bool) *UserService {
	s.locationDetails = locationDetails
	return s
}

func (s *UserService) PortfolioDetails(portfolioDetails bool) *UserService {
	s.portfolioDetails = portfolioDetails
	return s
}

func (s *UserService) PreferredDetails(preferredDetails bool) *UserService {
	s.preferredDetails = preferredDetails
	return s
}

func (s *UserService) BadgeDetails(badgeDetails bool) *UserService {
	s.badgeDetails = badgeDetails
	return s
}

func (s *UserService) Status(status bool) *UserService {
	s.status = status
	return s
}

func (s *UserService) Reputation(reputation bool) *UserService {
	s.reputation = reputation
	return s
}

func (s *UserService) EmployerReputation(employerReputation bool) *UserService {
	s.employerReputation = employerReputation
	return s
}

func (s *UserService) ReputationExtra(reputationExtra bool) *UserService {
	s.reputationExtra = reputationExtra
	return s
}

func (s *UserService) EmployerReputationExtra(employerReputationExtra bool) *UserService {
	s.employerReputationExtra = employerReputationExtra
	return s
}

func (s *UserService) CoverImage(coverImage bool) *UserService {
	s.coverImage = coverImage
	return s
}

func (s *UserService) PastCoverImage(pastCoverImage bool) *UserService {
	s.pastCoverImage = pastCoverImage
	return s
}

func (s *UserService) MobileTracking(mobileTracking bool) *UserService {
	s.mobileTracking = mobileTracking
	return s
}

func (s *UserService) BidQualityDetails(bidQualityDetails bool) *UserService {
	s.bidQualityDetails = bidQualityDetails
	return s
}

func (s *UserService) DepositMethods(depositMethods bool) *UserService {
	s.depositMethods = depositMethods
	return s
}

func (s *UserService) UserRecommendations(userRecommendations bool) *UserService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *UserService) MarketingMobileNumber(marketingMobileNumber bool) *UserService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *UserService) SanctionDetails(sanctionDetails bool) *UserService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *UserService) LimitedAccount(limitedAccount bool) *UserService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *UserService) CompletedUserRelevantJobCount(completedUserRelevantJobCount bool) *UserService {
	s.completedUserRelevantJobCount = completedUserRelevantJobCount
	return s
}

func (s *UserService) EquipmentGroupDetails(equipmentGroupDetails bool) *UserService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *UserService) JobRanks(jobRanks bool) *UserService {
	s.jobRanks = jobRanks
	return s
}

func (s *UserService) ShareholderDetails(shareholderDetails bool) *UserService {
	s.shareholderDetails = shareholderDetails
	return s
}

func (s *UserService) RisingStar(risingStar bool) *UserService {
	s.risingStar = risingStar
	return s
}

func (s *UserService) StaffDetails(staffDetails bool) *UserService {
	s.staffDetails = staffDetails
	return s
}

func (s *UserService) JobSeoDetails(jobSeoDetails bool) *UserService {
	s.jobSeoDetails = jobSeoDetails
	return s
}
