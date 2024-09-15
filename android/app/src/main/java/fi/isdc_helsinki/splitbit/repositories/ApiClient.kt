
import retrofit2.http.GET
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.POST



// Define your Retrofit API service
interface ApiService {
    @GET("groups/1/items")
    suspend fun getItems(): Response<List<Item>>
    @POST("login")
    suspend fun login(@Body req : AuthRequest): Response<AuthResponse>
    @POST("groups/1/items") // Unit in Kotlin signifies that no meaningful value is expected.
    suspend fun postItem(@Body item : Item): Response<Unit>
}

// Create a Retrofit instance
object Server {
    private val retrofit by lazy {
        Retrofit.Builder()
            .baseUrl("https://split-isdc.kuchta.dev")
            .addConverterFactory(GsonConverterFactory.create())
            .build()
    }

    val api: ApiService by lazy {
        retrofit.create(ApiService::class.java)
    }
}
