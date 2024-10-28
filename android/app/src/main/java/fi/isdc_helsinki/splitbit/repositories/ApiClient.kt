
import fi.isdc_helsinki.splitbit.client.apis.DefaultApi
import retrofit2.http.GET
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.POST



// Define your Retrofit API service


// Create a Retrofit instance
object Server {
    private val retrofit by lazy {
        Retrofit.Builder()
            .baseUrl("https://split-isdc.kuchta.dev")
            .addConverterFactory(GsonConverterFactory.create())
            .build()
    }

    val api: DefaultApi by lazy {
        retrofit.create(DefaultApi::class.java)
    }
}
