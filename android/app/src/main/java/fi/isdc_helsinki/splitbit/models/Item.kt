import androidx.annotation.Keep
import com.google.gson.annotations.SerializedName

// Define your data model
@Keep // THIS KEEP IS SUPER FUCKING IMPORTANT AS IT IS USED FOR REFLECTION, READ THIS DOC IT WASN'T REALLY DESCRIBED ANYWHERE
data class Item(
    @SerializedName("id")
    val id: Int,
    @SerializedName("name")
    val name: String,
    @SerializedName("price")
    val price: Float,
    @SerializedName("group_id")
    val group_id: Int,
    @SerializedName("author_id")
    val author_id: Int,
    @SerializedName("timestamp") val timestamp: Long)
