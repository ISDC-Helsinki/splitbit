package fi.isdc_helsinki.splitbit

import AuthRequest
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.ArrowBackIosNew
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.input.TextFieldValue
import androidx.lifecycle.viewmodel.compose.viewModel
import fi.isdc_helsinki.splitbit.client.models.UserCredentials

@Composable
fun LoginView(viewModel: LoginViewModel = viewModel()) {

    var step by remember { mutableStateOf(0) }
    var serverUrl by remember { mutableStateOf(TextFieldValue("")) }
    var user by remember { mutableStateOf(TextFieldValue("")) }
    var password by remember { mutableStateOf(TextFieldValue("")) }
    Scaffold { contentPadding ->
        Column(
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.Center,
            modifier = Modifier.fillMaxSize()
        ) {
            if (step == -1) {
                OutlinedTextField(serverUrl, { v -> serverUrl = v }, label = { Text("Server URL") })
                Text(AnnotatedString("Connection failed"))
                Button(onClick = { step = viewModel.checkServer(serverUrl.text) }) { Text("Next") }
            }
            else if (step == 0) {
                OutlinedTextField(serverUrl, { v -> serverUrl = v }, label = { Text("Server URL") })
                Button(onClick = { step = viewModel.checkServer(serverUrl.text) }) { Text("Next") }
            } else if (step > 0) {
                OutlinedTextField(user, { v -> user = v }, label = { Text("Login") })
                OutlinedTextField(password, { v -> password = v },
                    visualTransformation = PasswordVisualTransformation(),
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Password),
                    label = { Text("Password") })
                Row {
                    if (step > 0) {
                        IconButton(onClick = { step-- }) {
                            Icon(Icons.Filled.ArrowBackIosNew, contentDescription = "Add")
                        }
                    }
                    Button(onClick = {
                        viewModel.login(UserCredentials(user.text,password.text))
                        val id = g.encryptedSharedPreferences.getInt("userID", -1)
                        println("id $id")
                        if (id != -1) {
                            g.navController.navigate("group")
                        }
                    }) { Text("Login") }
                }
            }
            val id = g.encryptedSharedPreferences.getInt("userID", -1)
            println("id $id")
            if (id != -1) {
                g.navController.navigate("group")
            }
        }
    }

}