package freelancer

import (
	"context"
	"encoding/json"
	"net/http"
)

type UsersUsersService struct {
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

func (s *UsersUsersService) Do(ctx context.Context) (*ResponseUsers, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/user/0.1/users",
	}
	if len(s.users) > 0 {
		r.setParam("users", s.users)
	}
	if len(s.usernames) > 0 {
		r.setParam("usernames", s.usernames)
	}
	if s.avatar {
		r.setParam("avatar", s.avatar)
	}
	if s.countryDetails {
		r.setParam("country_details", s.countryDetails)
	}
	if s.profileDescription {
		r.setParam("profile_description", s.profileDescription)
	}
	if s.displayInfo {
		r.setParam("display_info", s.displayInfo)
	}
	if s.jobs {
		r.setParam("jobs", s.jobs)
	}
	if s.balanceDetails {
		r.setParam("balance_details", s.balanceDetails)
	}
	if s.qualificationDetails {
		r.setParam("qualification_details", s.qualificationDetails)
	}
	if s.membershipDetails {
		r.setParam("membership_details", s.membershipDetails)
	}
	if s.financialDetails {
		r.setParam("financial_details", s.financialDetails)
	}
	if s.locationDetails {
		r.setParam("location_details", s.locationDetails)
	}
	if s.portfolioDetails {
		r.setParam("portfolio_details", s.portfolioDetails)
	}
	if s.preferredDetails {
		r.setParam("preferred_details", s.preferredDetails)
	}
	if s.badgeDetails {
		r.setParam("badge_details", s.badgeDetails)
	}
	if s.status {
		r.setParam("status", s.status)
	}
	if s.reputation {
		r.setParam("reputation", s.reputation)
	}
	if s.employerReputation {
		r.setParam("employer_reputation", s.employerReputation)
	}
	if s.reputationExtra {
		r.setParam("reputation_extra", s.reputationExtra)
	}
	if s.employerReputationExtra {
		r.setParam("employer_reputation_extra", s.employerReputationExtra)
	}
	if s.coverImage {
		r.setParam("cover_image", s.coverImage)
	}
	if s.pastCoverImage {
		r.setParam("past_cover_image", s.pastCoverImage)
	}
	if s.mobileTracking {
		r.setParam("mobile_tracking", s.mobileTracking)
	}
	if s.bidQualityDetails {
		r.setParam("bid_quality_details", s.bidQualityDetails)
	}
	if s.depositMethods {
		r.setParam("deposit_methods", s.depositMethods)
	}
	if s.userRecommendations {
		r.setParam("user_recommendations", s.userRecommendations)
	}
	if s.marketingMobileNumber {
		r.setParam("marketing_mobile_number", s.marketingMobileNumber)
	}
	if s.sanctionDetails {
		r.setParam("sanction_details", s.sanctionDetails)
	}
	if s.limitedAccount {
		r.setParam("limited_account", s.limitedAccount)
	}
	if s.completedUserRelevantJobCount {
		r.setParam("completed_user_relevant_job_count", s.completedUserRelevantJobCount)
	}
	if s.equipmentGroupDetails {
		r.setParam("equipment_group_details", s.equipmentGroupDetails)
	}
	if s.jobRanks {
		r.setParam("job_ranks", s.jobRanks)
	}
	if s.jobSeoDetails {
		r.setParam("job_seo_details", s.jobSeoDetails)
	}
	if s.risingStar {
		r.setParam("rising_star", s.risingStar)
	}
	if s.shareholderDetails {
		r.setParam("shareholder_details", s.shareholderDetails)
	}
	if s.staffDetails {
		r.setParam("staff_details", s.staffDetails)
	}
	data, err := s.client.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := &ResponseUsers{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UsersUsersService) SetUsers(users []int64) *UsersUsersService {
	s.users = users
	return s
}

func (s *UsersUsersService) SetUsernames(usernames []string) *UsersUsersService {
	s.usernames = usernames
	return s
}

func (s *UsersUsersService) SetAvatar(avatar bool) *UsersUsersService {
	s.avatar = avatar
	return s
}

func (s *UsersUsersService) SetCountryDetails(countryDetails bool) *UsersUsersService {
	s.countryDetails = countryDetails
	return s
}

func (s *UsersUsersService) SetProfileDescription(profileDescription bool) *UsersUsersService {
	s.profileDescription = profileDescription
	return s
}

func (s *UsersUsersService) SetDisplayInfo(displayInfo bool) *UsersUsersService {
	s.displayInfo = displayInfo
	return s
}

func (s *UsersUsersService) SetJobs(jobs bool) *UsersUsersService {
	s.jobs = jobs
	return s
}

func (s *UsersUsersService) SetBalanceDetails(balanceDetails bool) *UsersUsersService {
	s.balanceDetails = balanceDetails
	return s
}

func (s *UsersUsersService) SetQualificationDetails(qualificationDetails bool) *UsersUsersService {
	s.qualificationDetails = qualificationDetails
	return s
}

func (s *UsersUsersService) SetMembershipDetails(membershipDetails bool) *UsersUsersService {
	s.membershipDetails = membershipDetails
	return s
}

func (s *UsersUsersService) SetFinancialDetails(financialDetails bool) *UsersUsersService {
	s.financialDetails = financialDetails
	return s
}

func (s *UsersUsersService) SetLocationDetails(locationDetails bool) *UsersUsersService {
	s.locationDetails = locationDetails
	return s
}

func (s *UsersUsersService) SetPortfolioDetails(portfolioDetails bool) *UsersUsersService {
	s.portfolioDetails = portfolioDetails
	return s
}

func (s *UsersUsersService) SetPreferredDetails(preferredDetails bool) *UsersUsersService {
	s.preferredDetails = preferredDetails
	return s
}

func (s *UsersUsersService) SetBadgeDetails(badgeDetails bool) *UsersUsersService {
	s.badgeDetails = badgeDetails
	return s
}

func (s *UsersUsersService) SetStatus(status bool) *UsersUsersService {
	s.status = status
	return s
}

func (s *UsersUsersService) SetReputation(reputation bool) *UsersUsersService {
	s.reputation = reputation
	return s
}

func (s *UsersUsersService) SetEmployerReputation(employerReputation bool) *UsersUsersService {
	s.employerReputation = employerReputation
	return s
}

func (s *UsersUsersService) SetReputationExtra(reputationExtra bool) *UsersUsersService {
	s.reputationExtra = reputationExtra
	return s
}

func (s *UsersUsersService) SetEmployerReputationExtra(employerReputationExtra bool) *UsersUsersService {
	s.employerReputationExtra = employerReputationExtra
	return s
}

func (s *UsersUsersService) SetCoverImage(coverImage bool) *UsersUsersService {
	s.coverImage = coverImage
	return s
}

func (s *UsersUsersService) SetPastCoverImage(pastCoverImage bool) *UsersUsersService {
	s.pastCoverImage = pastCoverImage
	return s
}

func (s *UsersUsersService) SetMobileTracking(mobileTracking bool) *UsersUsersService {
	s.mobileTracking = mobileTracking
	return s
}

func (s *UsersUsersService) SetBidQualityDetails(bidQualityDetails bool) *UsersUsersService {
	s.bidQualityDetails = bidQualityDetails
	return s
}

func (s *UsersUsersService) SetDepositMethods(depositMethods bool) *UsersUsersService {
	s.depositMethods = depositMethods
	return s
}

func (s *UsersUsersService) SetUserRecommendations(userRecommendations bool) *UsersUsersService {
	s.userRecommendations = userRecommendations
	return s
}

func (s *UsersUsersService) SetMarketingMobileNumber(marketingMobileNumber bool) *UsersUsersService {
	s.marketingMobileNumber = marketingMobileNumber
	return s
}

func (s *UsersUsersService) SetSanctionDetails(sanctionDetails bool) *UsersUsersService {
	s.sanctionDetails = sanctionDetails
	return s
}

func (s *UsersUsersService) SetLimitedAccount(limitedAccount bool) *UsersUsersService {
	s.limitedAccount = limitedAccount
	return s
}

func (s *UsersUsersService) SetCompletedUserRelevantJobCount(completedUserRelevantJobCount bool) *UsersUsersService {
	s.completedUserRelevantJobCount = completedUserRelevantJobCount
	return s
}

func (s *UsersUsersService) SetEquipmentGroupDetails(equipmentGroupDetails bool) *UsersUsersService {
	s.equipmentGroupDetails = equipmentGroupDetails
	return s
}

func (s *UsersUsersService) SetJobRanks(jobRanks bool) *UsersUsersService {
	s.jobRanks = jobRanks
	return s
}

func (s *UsersUsersService) SetShareholderDetails(shareholderDetails bool) *UsersUsersService {
	s.shareholderDetails = shareholderDetails
	return s
}

func (s *UsersUsersService) SetRisingStar(risingStar bool) *UsersUsersService {
	s.risingStar = risingStar
	return s
}

func (s *UsersUsersService) SetStaffDetails(staffDetails bool) *UsersUsersService {
	s.staffDetails = staffDetails
	return s
}

func (s *UsersUsersService) SetJobSeoDetails(jobSeoDetails bool) *UsersUsersService {
	s.jobSeoDetails = jobSeoDetails
	return s
}
