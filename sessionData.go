package irsdk

type SessionData struct {
	WeekendInfo WeekendInfo `yaml:",preserve"`
	SessionInfo SessionInfo  `yaml:",preserve"`
	QualifyResultsInfo QualifyResultsInfo  `yaml:",preserve"`
	CameraInfo CameraInfo `yaml:",preserve"`
	RadioInfo RadioInfo `yaml:",preserve"`
	DriverInfo DriverInfo `yaml:",preserve"`
	SplitTimeInfo SplitTimeInfo `yaml:",preserve"`
	//Purposefully did not implement car setup at this time.
}


type SessionInfo struct {
	Sessions []Session `yaml:",preserve"`
}

type Session struct {
	SessionNum SessionVariable `yaml:",preserve"`
	SessionLaps SessionVariable `yaml:",preserve"`
	SessionTime SessionVariable `yaml:",preserve"`
	SessionNumLapsToAvg SessionVariable `yaml:",preserve"`
	SessionType SessionVariable `yaml:",preserve"`
	SessionTrackRubberState SessionVariable `yaml:",preserve"`
	SessionName SessionVariable `yaml:",preserve"`
	SessionSubType SessionVariable `yaml:",preserve"`
	SessionSkipped SessionVariable `yaml:",preserve"`
	SessionRunGroupsUsed SessionVariable `yaml:",preserve"`
	ResultsPositions []ResultsPositions `yaml:",preserve"`
}

type ResultsPositions struct {
	Position SessionVariable `yaml:",preserve"`
	ClassPosition SessionVariable `yaml:",preserve"`
	CarIdx SessionVariable `yaml:",preserve"`
	Lap SessionVariable `yaml:",preserve"`
	Time SessionVariable `yaml:",preserve"`
	FastestLap SessionVariable `yaml:",preserve"`
	FastestTime SessionVariable `yaml:",preserve"`
	LastTime SessionVariable `yaml:",preserve"`
	LapsLed SessionVariable `yaml:",preserve"`
	LapsComplete SessionVariable `yaml:",preserve"`
	JokerLapsComplete SessionVariable `yaml:",preserve"`
	LapsDriven SessionVariable `yaml:",preserve"`
	Incidents SessionVariable `yaml:",preserve"`
	ReasonOutId SessionVariable `yaml:",preserve"`
	ReasonOutStr SessionVariable `yaml:",preserve"`
	ResultsFastestLap ResultsFastestLap `yaml:",preserve"`
	ResultsAverageLapTime SessionVariable `yaml:",preserve"`
	ResultsNumCautionFlags SessionVariable `yaml:",preserve"`
	ResultsNumCautionLaps SessionVariable `yaml:",preserve"`
	ResultsNumLeadChanges SessionVariable `yaml:",preserve"`
	ResultsLapsComplete SessionVariable `yaml:",preserve"`
	ResultsOfficial SessionVariable `yaml:",preserve"`
}

type ResultsFastestLap struct {
	CarIdx SessionVariable `yaml:",preserve"`
	FastestLap SessionVariable `yaml:",preserve"`
	FastestTime SessionVariable `yaml:",preserve"`
}

type WeekendInfo struct {
	TrackName SessionVariable `yaml:",preserve"`
	TrackID SessionVariable `yaml:",preserve"`
	TrackLength SessionVariable `yaml:",preserve"`
	TrackDisplayName SessionVariable `yaml:",preserve"`
	TrackDisplayShortName SessionVariable `yaml:",preserve"`
	TrackConfigName SessionVariable `yaml:",preserve"`
	TrackCity SessionVariable `yaml:",preserve"`
	TrackCountry SessionVariable `yaml:",preserve"`
	TrackAltitude SessionVariable `yaml:",preserve"`
	TrackLatitude SessionVariable `yaml:",preserve"`
	TrackLongitude SessionVariable `yaml:",preserve"`
	TrackNorthOffset SessionVariable `yaml:",preserve"`
	TrackNumTurns SessionVariable `yaml:",preserve"`
	TrackPitSpeedLimit SessionVariable `yaml:",preserve"`
	TrackType SessionVariable `yaml:",preserve"`
	TrackDirection SessionVariable `yaml:",preserve"`
	TrackWeatherType SessionVariable `yaml:",preserve"`
	TrackSkies SessionVariable `yaml:",preserve"`
	TrackSurfaceTemp Temperature `yaml:",preserve"`
	TrackAirTemp Temperature `yaml:",preserve"`
	TrackAirPressure SessionVariable `yaml:",preserve"`
	TrackWindVel SessionVariable `yaml:",preserve"`
	TrackWindDir SessionVariable `yaml:",preserve"`
	TrackRelativeHumidity SessionVariable `yaml:",preserve"`
	TrackFogLevel SessionVariable `yaml:",preserve"`
	TrackCleanup SessionVariable `yaml:",preserve"`
	TrackDynamicTrack SessionVariable `yaml:",preserve"`
	SeriesID SessionVariable `yaml:",preserve"`
	SeasonID SessionVariable `yaml:",preserve"`
	SessionID SessionVariable `yaml:",preserve"`
	SubSessionID SessionVariable `yaml:",preserve"`
	LeagueID SessionVariable `yaml:",preserve"`
	Official SessionVariable `yaml:",preserve"`
	RaceWeek SessionVariable `yaml:",preserve"`
	EventType SessionVariable `yaml:",preserve"`
	Category SessionVariable `yaml:",preserve"`
	SimMode SessionVariable `yaml:",preserve"`
	TeamRacing SessionVariable `yaml:",preserve"`
	MinDrivers SessionVariable `yaml:",preserve"`
	MaxDrivers SessionVariable `yaml:",preserve"`
	DCRuleSet SessionVariable `yaml:",preserve"`
	QualifierMustStartRace SessionVariable `yaml:",preserve"`
	NumCarClasses SessionVariable `yaml:",preserve"`
	NumCarTypes SessionVariable `yaml:",preserve"`
	HeatRacing SessionVariable `yaml:",preserve"`
	WeekendOptions WeekendOptions `yaml:",preserve"`
	TelemetryOptions TelemetryOptions  `yaml:",preserve"`
}


type WeekendOptions struct {
	NumStarters SessionVariable `yaml:",preserve"`
	StartingGrid SessionVariable `yaml:",preserve"`
	QualifyScoring SessionVariable `yaml:",preserve"`
	CourseCautions SessionVariable `yaml:",preserve"`
	StandingStart SessionVariable `yaml:",preserve"`
	Restarts SessionVariable `yaml:",preserve"`
	WeatherType SessionVariable `yaml:",preserve"`
	Skies SessionVariable `yaml:",preserve"`
	WindDirection SessionVariable `yaml:",preserve"`
	WindSpeed SessionVariable `yaml:",preserve"`
	WeatherTemp SessionVariable `yaml:",preserve"`
	RelativeHumidity SessionVariable `yaml:",preserve"`
	FogLevel SessionVariable `yaml:",preserve"`
	TimeOfDay SessionVariable `yaml:",preserve"`
	Date SessionVariable `yaml:",preserve"`
	EarthRotationSpeedupFactor SessionVariable `yaml:",preserve"`
	Unofficial SessionVariable `yaml:",preserve"`
	CommercialMode SessionVariable `yaml:",preserve"`
	NightMode SessionVariable `yaml:",preserve"`
	IsFixedSetup SessionVariable `yaml:",preserve"`
	StrictLapsChecking SessionVariable `yaml:",preserve"`
	HasOpenRegistration SessionVariable `yaml:",preserve"`
	HardcoreLevel SessionVariable `yaml:",preserve"`
	NumJokerLaps SessionVariable `yaml:",preserve"`
	IncidentLimit SessionVariable `yaml:",preserve"`
}

type TelemetryOptions struct {
	TelemetryDiskFile SessionVariable `yaml:",preserve"`
}


type QualifyResultsInfo struct {
	Results []QualifyResult `yaml:",preserve"`
}

type QualifyResult struct {
	Position      SessionVariable `yaml:",preserve"`
	ClassPosition SessionVariable `yaml:",preserve"`
	CarIdx        SessionVariable `yaml:",preserve"`
	FastestLap    SessionVariable `yaml:",preserve"`
	FastestTime   SessionVariable `yaml:",preserve"`
}

type CameraInfo struct {
	Groups []CameraGroup `yaml:",preserve"`
}

type CameraGroup struct {
	GroupNum SessionVariable `yaml:",preserve"`
	GroupName SessionVariable `yaml:",preserve"`
	Cameras []Cameras `yaml:",preserve"`
}

type Cameras struct {
	CameraNum SessionVariable `yaml:",preserve"`
	CameraName SessionVariable `yaml:",preserve"`
}

type RadioInfo struct {
	SelectedRadioNum SessionVariable `yaml:",preserve"`
	Radios []Radio
}

type Radio struct {
	RadioNum SessionVariable `yaml:",preserve"`
	HopCount SessionVariable `yaml:",preserve"`
	NumFrequencies SessionVariable `yaml:",preserve"`
	TunedToFrequencyNum SessionVariable `yaml:",preserve"`
	ScanningIsOn SessionVariable `yaml:",preserve"`
	Frequencies []RadioFrequency
}

type RadioFrequency struct {
	FrequencyNum SessionVariable `yaml:",preserve"`
	FrequencyName SessionVariable `yaml:",preserve"`
	Priority SessionVariable `yaml:",preserve"`
	CarIdx SessionVariable `yaml:",preserve"`
	EntryIdx SessionVariable `yaml:",preserve"`
	ClubID SessionVariable `yaml:",preserve"`
	CanScan SessionVariable `yaml:",preserve"`
	CanSquawk SessionVariable `yaml:",preserve"`
	Muted SessionVariable `yaml:",preserve"`
	IsMutable SessionVariable `yaml:",preserve"`
	IsDeletable SessionVariable `yaml:",preserve"`
}

type DriverInfo struct {
	DriverCarIdx SessionVariable `yaml:",preserve"`
	DriverUserID SessionVariable `yaml:",preserve"`
	PaceCarIdx SessionVariable `yaml:",preserve"`
	DriverHeadPosX SessionVariable `yaml:",preserve"`
	DriverHeadPosY SessionVariable `yaml:",preserve"`
	DriverHeadPosZ SessionVariable `yaml:",preserve"`
	DriverCarIdleRPM SessionVariable `yaml:",preserve"`
	DriverCarRedLine SessionVariable `yaml:",preserve"`
	DriverCarEngCylinderCount SessionVariable `yaml:",preserve"`
	DriverCarFuelKgPerLtr SessionVariable `yaml:",preserve"`
	DriverCarFuelMaxLtr SessionVariable `yaml:",preserve"`
	DriverCarMaxFuelPct SessionVariable `yaml:",preserve"`
	DriverCarSLFirstRPM SessionVariable `yaml:",preserve"`
	DriverCarSLShiftRPM SessionVariable `yaml:",preserve"`
	DriverCarSLLastRPM SessionVariable `yaml:",preserve"`
	DriverCarSLBlinkRPM SessionVariable `yaml:",preserve"`
	DriverPitTrkPct SessionVariable `yaml:",preserve"`
	DriverCarEstLapTime SessionVariable `yaml:",preserve"`
	DriverSetupName SessionVariable `yaml:",preserve"`
	DriverSetupIsModified SessionVariable `yaml:",preserve"`
	DriverSetupLoadTypeName SessionVariable `yaml:",preserve"`
	DriverSetupPassedTech SessionVariable `yaml:",preserve"`
	DriverIncidentCount SessionVariable `yaml:",preserve"`
	Drivers []Driver `yaml:",preserve"`
}

type Driver struct {
	CarIdx SessionVariable `yaml:",preserve"`
	UserName SessionVariable `yaml:",preserve"`
	AbbrevName SessionVariable `yaml:",preserve"`
	Initials SessionVariable `yaml:",preserve"`
	UserID SessionVariable `yaml:",preserve"`
	TeamID SessionVariable `yaml:",preserve"`
	TeamName SessionVariable `yaml:",preserve"`
	CarNumber SessionVariable `yaml:",preserve"`
	CarNumberRaw SessionVariable `yaml:",preserve"`
	CarPath SessionVariable `yaml:",preserve"`
	CarClassID SessionVariable `yaml:",preserve"`
	CarID SessionVariable `yaml:",preserve"`
	CarIsPaceCar SessionVariable `yaml:",preserve"`
	CarIsAI SessionVariable `yaml:",preserve"`
	CarScreenName SessionVariable `yaml:",preserve"`
	CarScreenNameShort SessionVariable `yaml:",preserve"`
	CarClassShortName SessionVariable `yaml:",preserve"`
	CarClassRelSpeed SessionVariable `yaml:",preserve"`
	CarClassLicenseLevel SessionVariable `yaml:",preserve"`
	CarClassMaxFuelPct SessionVariable `yaml:",preserve"`
	CarClassWeightPenalty SessionVariable `yaml:",preserve"`
	CarClassColor SessionVariable `yaml:",preserve"`
	IRating SessionVariable `yaml:",preserve"`
	LicLevel SessionVariable `yaml:",preserve"`
	LicSubLevel SessionVariable `yaml:",preserve"`
	LicString SessionVariable `yaml:",preserve"`
	LicColor SessionVariable `yaml:",preserve"`
	IsSpectator SessionVariable `yaml:",preserve"`
	CarDesignStr SessionVariable `yaml:",preserve"`
	HelmetDesignStr SessionVariable `yaml:",preserve"`
	SuitDesignStr SessionVariable `yaml:",preserve"`
	CarNumberDesignStr SessionVariable `yaml:",preserve"`
	CarSponsor_1 SessionVariable `yaml:",preserve"`
	CarSponsor_2 SessionVariable `yaml:",preserve"`
	ClubName SessionVariable `yaml:",preserve"`
	DivisionName SessionVariable `yaml:",preserve"`
	CurDriverIncidentCount SessionVariable `yaml:",preserve"`
	TeamIncidentCount SessionVariable `yaml:",preserve"`
}

type SplitTimeInfo struct {
	Sectors []Sector `yaml:",preserve"`
}

type Sector struct {
	SectorNum SessionVariable `yaml:",preserve"`
	SectorStartPct SessionVariable `yaml:",preserve"`
}
