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

var FixedAllowedRoles = [2]string{"superadmin", "admin"}

const (
	WATER_METER_PATH = "assets/water-meter/"
)

const (
	MINIO_BUCKET_ORIGINAL = "original-images"
	MINIO_BUCKET_CROPPED  = "cropped-images"
	MINIO_BUCKET_FILE     = "files"
)
