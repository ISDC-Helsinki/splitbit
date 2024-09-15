package fi.isdc_helsinki.splitbit
import android.content.Context
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.rememberNavController

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge

import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.platform.LocalContext
import androidx.navigation.NavHostController
import androidx.navigation.compose.composable
import com.example.splitbit.ui.theme.SplitBitTheme
import fi.isdc_helsinki.splitbit.ui.theme.GroupView
import java.util.Scanner

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            SplitBitTheme {
                SplitBit()
            }
        }
    }
}
// Global singleton
object g {
    lateinit var navController: NavHostController
}
@Composable
fun SplitBit() {
    val navController = rememberNavController()
    g.navController = navController
    SplitBitNavHost(navController)
}

@Composable
fun SplitBitNavHost(navController: NavHostController) {
    val context = LocalContext.current
    val sharedPreferences = context.getSharedPreferences("app_prefs", Context.MODE_PRIVATE)
    val isLoggedIn = sharedPreferences.getBoolean("isLoggedIn", false)
    val startDestination = if (isLoggedIn) "group" else "group" // TODO login

    NavHost(navController, startDestination = startDestination) {
        composable("group") { GroupView() }
        composable("login") { LoginView() }
        composable("add") { ExpenseAddView() }
        composable("scanner") { ScannerView() }
    }
}
