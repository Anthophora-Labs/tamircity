package db

import "gorm.io/gorm"

type ExpertisePhoneInfo struct {
	gorm.Model
	ReservationId             uint
	Reservation               *Reservation `gorm:"foreignkey:ReservationId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Invoice                   bool         //Fatura var mi
	Box                       bool         //Kutu
	GuaranteeTerm             int          //1-24
	Color                     string
	IMEIRegistration          string
	EDevletRegistration       string
	ScreenSize                float32
	ScreenTechnology          string
	ScreenResolution          string
	ScratchResistance         bool
	CPUModel                  string
	CPUFrequency              int
	Ram                       int
	OSType                    string
	OSTypeVersion             string
	CPUCoreNumber             int
	CameraResolution          int
	FrontCameraResolution     int
	VideoRecordResolution     string
	VideoFPS                  int
	FaceRecognition           bool
	SlowMotionVideo           bool
	CameraAI                  bool
	Timer                     bool
	AutomaticFocus            bool
	GeographicLocation        bool
	VoiceControl              bool
	InternalStorage           int
	ExternalStorage           int
	MaxExternalStorage        int
	BatteryType               string
	BatteryCapacity           int
	BatteryWirelessCharge     bool
	BatteryFastCharge         bool
	BatteryWirelessFastCharge bool
	BatteryDetachable         bool
	WiFiFrequency             string
	NFC                       bool
	G5Support                 bool
	ReleaseYear               int
	ResistanceofWater         bool
	ResistanceofDust          bool
	Fingerprint               bool
	DoubleSim                 bool
	AnTuTuScore               int

	IsScreenHasBrokenProblem        bool
	IsScreenHasObscurationProblem   bool
	IsTouchScreenHasProblem         bool
	IsScreenHasDeadPixelPixel       bool
	IsDeviceHasCaseProblem          bool
	IsDeviceHasCoverProblem         bool
	IsDeviceHaveCamerasProblem      bool
	IsDeviceHaveSpeakerProblem      bool
	IsDeviceHasHighHeatProblem      bool
	IsDeviceHasChargeSocketProblem  bool
	IsDeviceHasPowerButtonProblem   bool
	IsDeviceHasOpenedCaseProblem    bool
	IsDeviceHasSideButtonProblem    bool
	IsDeviceHasFreezingProblem      bool
	IsDeviceHasBluetoothProblem     bool
	IsDeviceHasWiFiProblem          bool
	IsDeviceHasMicrophoneProblem    bool
	IsDeviceHasCellularProblem      bool
	IsDeviceHasSoundTransferProblem bool
}
