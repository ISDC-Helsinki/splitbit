package fi.isdc_helsinki.splitbit.ui.theme

import Item
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.material3.TopAppBarDefaults.topAppBarColors
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.input.nestedscroll.nestedScroll
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import fi.isdc_helsinki.splitbit.ExpenseListItem
import fi.isdc_helsinki.splitbit.client.apis.DefaultApi
import fi.isdc_helsinki.splitbit.g
import fi.isdc_helsinki.splitbit.group.GroupViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

@OptIn(ExperimentalMaterial3Api::class)
@Preview(showBackground = true)
@Composable
fun GroupView(viewModel: GroupViewModel = viewModel()) {
    val itemList by viewModel.items.collectAsState()
    var presses by remember { mutableIntStateOf(0) }
    val sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true)
    val scope = rememberCoroutineScope()
    var showBottomSheet by remember { mutableStateOf(false) }
    val scrollBehavior =
        TopAppBarDefaults.exitUntilCollapsedScrollBehavior(rememberTopAppBarState())
    var selectedItem by remember { mutableIntStateOf(0) }

    LaunchedEffect(Unit) {
        viewModel.fetchItems()
    }

    Scaffold(
        modifier = Modifier.nestedScroll(scrollBehavior.nestedScrollConnection),
        topBar = {
            LargeTopAppBar(
                colors = topAppBarColors(
                    containerColor = MaterialTheme.colorScheme.primaryContainer,
                    titleContentColor = MaterialTheme.colorScheme.primary,
                ),
                title = {
                    Text("Home Group")
                },
                scrollBehavior = scrollBehavior
            )
        },
        bottomBar = {
            NavigationBar {
                NavigationBarItem(
                    icon = { Icon(Icons.Filled.Groups, contentDescription = "Groups") },
                    label = { Text("Groups")},
                    selected = selectedItem == 0,
                    onClick = { selectedItem = 0 }
                )
            }
        },
        floatingActionButton = {
            ExtendedFloatingActionButton(
                onClick = { g.navController.navigate("add") },
                icon = { Icon(Icons.Filled.Add, "Add Expense") },
                text = { Text(text = "Add Expense") },
            )
        }
    ) { innerPadding ->
        if (showBottomSheet) {
            ModalBottomSheet(
                modifier = Modifier.fillMaxHeight(0.8f),
                onDismissRequest = {
                    showBottomSheet = false
                },
                sheetState = sheetState

            ) {
                Column(modifier = Modifier.padding(16.dp))
                {
                    // Sheet content
                }
            }
        }
        Column(
            modifier = Modifier.padding(innerPadding),
        ) {
            Column(modifier = Modifier.padding(10.dp)) {
                Text("You owe: 512")
                Button(onClick = { showBottomSheet = true }) {
                    Text("Settle")
                }
                Button(onClick = { g.navController.navigate("login")}) {
                    Text("TestButton")
                }
            }
            if(itemList.isEmpty()) {
                Text("Loading")
            } else {
                LazyColumn {
                    itemsIndexed(itemList) { idx, item ->
                        ExpenseListItem(
                            icon = {
                                Icon(
                                    imageVector = Icons.Filled.LocalGroceryStore, // Material Icon
                                    contentDescription = "Groceries",
                                    modifier = Modifier.size(24.dp)
                                )
                            },
                            title = item.name,
                            amount = item.price,
                            payerInfo = "Mariusz paid \$"+item.price,
                            timestamp = item.timestamp
                        )
                    }
                }
            }
        }
    }
}

