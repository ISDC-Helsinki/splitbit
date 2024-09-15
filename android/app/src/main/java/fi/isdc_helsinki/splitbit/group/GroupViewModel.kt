package fi.isdc_helsinki.splitbit.group

import Item
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import Server


class GroupViewModel : ViewModel() {
    private val _items = MutableStateFlow<List<Item>>(emptyList())
    val items: StateFlow<List<Item>> = _items

    // Make the API call in a coroutine
    // This should be abstracted to a repository
    fun fetchItems() {
        viewModelScope.launch {
            try {
                val response = Server.api.getItems()
                if (response.isSuccessful) {
                    response.body()?.let {
                        _items.value = it
                    }
                }
            } catch (e: Exception) {
                // Handle error here
            }
        }
    }
    fun addItem(newItem: Item) {
        viewModelScope.launch {
            try {
                val response = Server.api.postItem(newItem)
                if (response.isSuccessful) {
                    response.body()?.let { _ ->
                        // Update local state if necessary
                        _items.value = _items.value + newItem
                    }
                }
            } catch (e: Exception) {
                // Handle error here
            }
        }
    }
}