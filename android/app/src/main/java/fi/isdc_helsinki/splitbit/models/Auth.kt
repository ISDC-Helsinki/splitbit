import androidx.annotation.Keep
import com.google.gson.annotations.SerializedName

@Keep
data class AuthRequest(
    @SerializedName("username")
    val username: String,
    @SerializedName("password")
    val password: String
)

@Keep
data class AuthResponse(
    @SerializedName("token")
    val token: String)

