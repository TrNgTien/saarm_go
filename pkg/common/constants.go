package common

const (
	L50       = 50
	L100      = 100
	L500      = 500
	L1000     = 1_000
	L5000     = 5_000
	JwtBearer = "Bearer"
	Jwt2FA    = "2FA"
	JwtOauth2 = "Oauth2"
)

var FixedAllowedRoles = []string{"superadmin", "admin"}

const (
	SUPER_ADMIN_ROLE = "superadmin"
	ADMIN_ROLE       = "admin"
	GUEST_ROLE       = "guest"
	LANDLORD_ROLE    = "landlord"
)

const (
	DELETED_STATUS   = "000_DELETED"
	ACTIVATED_STATUS = "100_ACTIVATED"
	DEACTIVED_STATUS = "200_DEACTIVED"
	ARCHIVE_STATUS   = "300_ARCHIVE"
)

const (
	WATER_METER_PATH = "assets/water-meter/"
)

const (
	MINIO_BUCKET_ORIGINAL = "original-images"
	MINIO_BUCKET_CROPPED  = "cropped-images"
	MINIO_BUCKET_FILE     = "files"
)

const (
	USER_PATH      = "/users"
	AUTH_PATH      = "/auth"
	ROOM_PATH      = "/rooms"
	ROLE_PATH      = "/roles"
	APARTMENT_PATH = "/apartments"
	CONFIG_PATH    = "/configs"
)
