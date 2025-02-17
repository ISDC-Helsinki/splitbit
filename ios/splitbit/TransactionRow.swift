
import SwiftUI

struct TransactionRow: View {
    var date: String
    var icon: String
    var description: String
    var paidInfo: String
    var amount: String
    
    var body: some View {
        HStack {
            // Date and Icon
            VStack(alignment: .leading) {
                Text(date)
                    .font(.headline)
                Image(systemName: icon)
                    .font(.largeTitle)
                    .padding(.top, 5)
            }
            
            Spacer()
            
            // Description and Paid Info
            VStack(alignment: .leading) {
                Text(description)
                    .font(.headline)
                Text(paidInfo)
                    .font(.subheadline)
                    .foregroundColor(.gray)
            }
            
            Spacer()
            
            // Borrowed Info
            VStack(alignment: .trailing) {
                Text("You borrowed")
                    .font(.subheadline)
                    .foregroundColor(.gray)
                Text(amount)
                    .font(.title)
                    .fontWeight(.bold)
                    .foregroundColor(.blue)
            }
        }
        .padding()
        .background(Color.black.opacity(0.9))
        .cornerRadius(10)
        .foregroundColor(.white)
    }
}

struct TransactionRow_Previews: PreviewProvider {
    static var previews: some View {
        TransactionRow(date: "Sept 31", icon: "car.fill", description: "Taxi drive", paidInfo: "Amalia paid $63", amount: "32.30$")
            .previewLayout(.sizeThatFits)
            .padding()
            .background(Color.black.edgesIgnoringSafeArea(.all))
    }
}
