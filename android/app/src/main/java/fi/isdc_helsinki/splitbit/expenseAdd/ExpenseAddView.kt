package fi.isdc_helsinki.splitbit


import Item
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.imePadding
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.*
import androidx.compose.material.icons.outlined.LocationOn
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.TextFieldValue
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import java.time.Instant

@Composable
fun ExpenseAdd(
    product: TextFieldValue,
    onProductChage: (TextFieldValue) -> Unit,
    amount: TextFieldValue,
    onAmountChange: (TextFieldValue) -> Unit
) {
    var paidBy by remember { mutableStateOf("you") }
    var splitMethod by remember { mutableStateOf("equally") }

    Column(
        modifier = Modifier
            .fillMaxWidth()
            .padding(16.dp)
    ) {
        // Row with Icon and Title
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.fillMaxWidth()
        ) {
            // Product Icon
            Box(
                modifier = Modifier
                    .size(40.dp)
                    .background(MaterialTheme.colorScheme.primaryContainer, shape = CircleShape),
                contentAlignment = Alignment.Center
            ) {
                Icon(Icons.Default.ShoppingCart, contentDescription = "Add")
            }

            Spacer(modifier = Modifier.width(8.dp))

            // Amount Input Field
            OutlinedTextField(
                value = product,
                onValueChange = { newValue -> onProductChage(newValue) },
                textStyle = MaterialTheme.typography.bodyLarge.copy(fontSize = 24.sp),
                modifier = Modifier.fillMaxWidth()
            )
        }

        Spacer(modifier = Modifier.height(8.dp))

        Divider()

        Spacer(modifier = Modifier.height(8.dp))

        // Row with Currency Icon and TextField
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.fillMaxWidth()
        ) {
            // Currency Icon
            Box(
                modifier = Modifier
                    .size(40.dp)
                    .background(MaterialTheme.colorScheme.primaryContainer, shape = CircleShape),
                contentAlignment = Alignment.Center
            ) {
                Text(
                    text = "$",
                    style = MaterialTheme.typography.bodyLarge.copy(fontSize = 20.sp),
                    fontWeight = FontWeight.Bold
                )
            }

            Spacer(modifier = Modifier.width(8.dp))

            // Amount Input Field
            OutlinedTextField(
                value = amount,
                onValueChange = { newValue -> onAmountChange(newValue) },
                textStyle = MaterialTheme.typography.bodyLarge.copy(fontSize = 24.sp),
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                modifier = Modifier.fillMaxWidth()
            )
        }

        Spacer(modifier = Modifier.height(16.dp))

        // Paid by and Split Method
        Row(
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.Center,
            modifier = Modifier.fillMaxWidth()

        ) {
            Text(text = "Paid by", style = MaterialTheme.typography.bodyLarge)

            Spacer(modifier = Modifier.width(8.dp))

            // Paid by Chip/Button
            Button(
                onClick = { /* Handle click */ },
                modifier = Modifier
                    .padding(horizontal = 4.dp)
            ) {
                Text(text = paidBy)
            }

            Spacer(modifier = Modifier.width(8.dp))

            Text(text = "and split")

            Spacer(modifier = Modifier.width(8.dp))

            // Split method Chip/Button
            Button(
                onClick = { /* Handle click */ },
                modifier = Modifier
                    .padding(horizontal = 4.dp)
            ) {
                Text(text = splitMethod, style = MaterialTheme.typography.bodyMedium)
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ExpenseAddView(viewModel: ExpenseAddViewModel = viewModel()) {
    var selected by remember { mutableStateOf(false) }
    var amount by remember { mutableStateOf(TextFieldValue("")) }
    var product by remember { mutableStateOf(TextFieldValue("")) }

    Scaffold(
        topBar = {
            TopAppBar(
                navigationIcon = {
                    IconButton(onClick = { g.navController.popBackStack()}) {
                        Icon(
                            imageVector = Icons.AutoMirrored.Filled.ArrowBack,
                            contentDescription = "Localized description"
                        )
                    }
                },
                title = {
                    Text("Adding to GayHQ")
                },
//                actions = {
//                    IconButton(onClick = { /* do something */ }) {
//                        Icon(
//                            imageVector = Icons.Filled.Check,
//                            contentDescription = "Localized description"
//                        )
//                    }
//                }
            )
        },
        bottomBar = {
            BottomAppBar(
                modifier = Modifier.imePadding(),
//                containerColor = MaterialTheme.colorScheme.primaryContainer,
//                contentColor = MaterialTheme.colorScheme.primary,
                actions = {
                    IconButton(onClick = { /* do something */ }) {
                        Icon(
                            Icons.Filled.Image,
                            contentDescription = "Localized description",
                        )
                    }
                    IconButton(onClick = { /* do something */ }) {
                        Icon(
                            Icons.Filled.DateRange,
                            contentDescription = "Localized description",
                        )
                    }
                    IconButton(onClick = { /* do something */ }) {
                        Icon(
                            Icons.Outlined.LocationOn,
                            contentDescription = "Localized description"
                        )
                    }
                    IconButton(onClick = { /* do something */ }) {
                        Icon(
                            Icons.Filled.Description,
                            contentDescription = "Localized description",
                        )
                    }
                },
                floatingActionButton = {
                    FloatingActionButton(
                        onClick = {
                            val newItem = Item(id=1, name = product.text , price = amount.text.toFloat(), group_id = 1, author_id = 1, timestamp = Instant.now().epochSecond)
                            viewModel.addItem(newItem)
                            g.navController.popBackStack()

                        },
                        containerColor = BottomAppBarDefaults.bottomAppBarFabColor,
                        elevation = FloatingActionButtonDefaults.bottomAppBarFabElevation()
                    ) {
                        Icon(Icons.Filled.Check, "Add Item")
                    }
                }
            )
        }
    ) { innerPadding ->
        Column(
            modifier = Modifier
                .padding(innerPadding)
                .fillMaxHeight()
        ) {

            Row(
                verticalAlignment = Alignment.CenterVertically,
                horizontalArrangement = Arrangement.Center,
                modifier = Modifier.fillMaxWidth()
            )
            {
                Text("With you and")
                Spacer(modifier = Modifier.width(10.dp))
                FilterChip(
                    onClick = { selected = !selected },
                    label = {
                        Text("Whole Group")
                    },
                    selected = selected,
                    leadingIcon = if (selected) {
                        {
                            Icon(
                                imageVector = Icons.Filled.Done,
                                contentDescription = "Done icon",
                                modifier = Modifier.size(FilterChipDefaults.IconSize)
                            )
                        }
                    } else {
                        null
                    },
                )
            }
            HorizontalDivider()
            Spacer(modifier = Modifier.height(30.dp))
            ExpenseAdd(product, { newText -> product = newText}, amount, {newAmount -> amount = newAmount})
        }
    }
}
