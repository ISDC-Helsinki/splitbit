//
//  Homegroup.swift
//  splitbit
//
//  Created by Anastasia on 28/10/2024.
//

import SwiftUI

struct Homegroup: View {
    var body: some View {
        
        VStack(alignment: .leading) {
            Text("You owe 450 zl overall")
            Button("Settle") {
                
            }.buttonStyle(.borderedProminent)
            
            
            List(0..<20) { item in
                TransactionRow(date: "Sept 31", icon: "car.fill", description: "Taxi drive", paidInfo: "Amalia paid $63", amount: "32.30$")
            }
        }
        .padding(.leading)
        .navigationTitle("Home group")
    }
}
