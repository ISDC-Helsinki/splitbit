package fi.isdc_helsinki.splitbit

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.compose.ui.text.font.FontWeight
import java.time.*
import java.time.format.DateTimeFormatter


fun unixTimeToFormattedDate(unixTime: Long): String {
    // Convert seconds to Instant
    val instant = Instant.ofEpochSecond(unixTime)
    val zonedDateTime = ZonedDateTime.ofInstant(instant, ZoneId.systemDefault())
    return zonedDateTime.format(DateTimeFormatter.ofPattern("MMM dd"))
}

@Composable
fun ExpenseListItem(
    icon: @Composable () -> Unit, // Now using a composable function for the icon
    title: String,
    payerInfo: String,
    amount: Float,
    timestamp: Long
) {
    val date = unixTimeToFormattedDate(timestamp)
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .padding(8.dp),
        verticalAlignment = Alignment.CenterVertically
    ) {
        // Date column
        Column(
            modifier = Modifier.width(50.dp),
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Text(
                text = date.split(" ")[0], // Day part
                style = MaterialTheme.typography.bodyMedium.copy(fontWeight = FontWeight.Bold)
            )
            Text(
                text = date.split(" ")[1], // Month part
                style = MaterialTheme.typography.bodySmall
            )
        }

        Spacer(modifier = Modifier.width(8.dp))

        // Icon with circular background
        Box(
            modifier = Modifier
                .size(40.dp)
                .background(color = MaterialTheme.colorScheme.primaryContainer, shape = CircleShape),
            contentAlignment = Alignment.Center
        ) {
            icon() // Invoke the icon composable
        }

        Spacer(modifier = Modifier.width(8.dp))

        // Title and payer info
        Column(modifier = Modifier.weight(1f)) {
            Text(
                text = title,
                style = MaterialTheme.typography.bodyLarge.copy(fontWeight = FontWeight.Bold)
            )
            Text(
                text = payerInfo,
                style = MaterialTheme.typography.bodySmall
            )
        }
        if(amount > 0) {
            Column(horizontalAlignment = Alignment.End) {
                Text(
                    text = "You lent",
                    color = MaterialTheme.colorScheme.tertiary,
                    style = MaterialTheme.typography.bodySmall//.copy(color = Color(0xFFFF6E40))
                )
                Text(
                    text = amount.toString() + "$",
                    color = MaterialTheme.colorScheme.tertiary,
                    style = MaterialTheme.typography.bodyLarge.copy(
                        fontWeight = FontWeight.Bold
                    )
                )
            }
        }
        else {
            Column(horizontalAlignment = Alignment.End) {
                Text(
                    text = "You borrowed",
                    color = MaterialTheme.colorScheme.error,
                    style = MaterialTheme.typography.bodySmall//.copy(color = Color(0xFFFF6E40))
                )
                Text(
                    text =  amount.toString() + "$" ,
                    color = MaterialTheme.colorScheme.error,
                    style = MaterialTheme.typography.bodyLarge.copy(
                        fontWeight = FontWeight.Bold
                    )
                )
            }
        }
        // Amount and label
    }
}
