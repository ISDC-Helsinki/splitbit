//
//  Addpay.swift
//  splitbit
//
//  Created by Anastasia on 28/10/2024.
//

import SwiftUI


struct Addpay: View {
    
    @State private var date = Date()
    @State private var item: String = ""
    @State private var price: String = ""
    @State private var tagindex = 0
    
    
    var body: some View {
        
        VStack(alignment: .leading){
            
            Text("Add new item")
            
            TextField("Item", text: $item)
                .padding(15) // Internal padding inside the text field
                .background(Color("Primary colour")) // Background color for the text field
                .cornerRadius(10) // Corner radius for rounded edges
                .overlay( // Overlay the border using the same corner radius
                    RoundedRectangle(cornerRadius: 10)
                        .stroke(Color.black, lineWidth: 2) // Reference custom colour from assets
                )
                .padding(10)
            
            Text("Price")
            
            
            TextField("€ £ $", text: $price)
                .keyboardType(.decimalPad)
                .padding(15)
                .background(Color("Primary colour"))
                .cornerRadius(10)
                .overlay(
                    RoundedRectangle(cornerRadius: 10)
                        .stroke(Color.black, lineWidth: 2)
                )
                .padding(10)
                .padding(.horizontal, 15)
            
            
            DatePicker(
                "Date",
                selection: $date,
                displayedComponents: [.date]
            )
            .padding(.vertical)
            
            
            
            HStack{
                
                Text("Tag")
                
                Picker("Tag", selection: $tagindex, content: {
                    Text("Travel").tag(0)
                    Text("Food").tag(1)
                    Text("House").tag(2)
                })
                .padding(10)
                .background(Color.white)
                .cornerRadius(10)
                .overlay(
                    RoundedRectangle(cornerRadius: 10)
                        .stroke(Color.black, lineWidth: 2))
                .padding(.horizontal,80)

                Group {
                    if tagindex == 0 {
                        Image(systemName: "airplane")
                            .font(.system(size: 30))
                    } else if tagindex == 1 {
                        Image(systemName: "fork.knife")
                            .font(.system(size: 30))
                    } else if tagindex == 2 {
                        Image(systemName: "house.fill")
                            .font(.system(size: 30))
                    }
                }
                Spacer()
            }
            // end of Hstack
            
            
            Button(
                action: {UIApplication.shared.sendAction(#selector(UIResponder.resignFirstResponder), to: nil,
                                                         from: nil, for: nil)}) {
                    Label("Done", systemImage: "chevron.right.circle")
                        .padding()
                        .background(Color("Primary colour"))
                        .foregroundColor(.black)
                        .cornerRadius(10)
                        .overlay(RoundedRectangle(cornerRadius: 10)
                        .stroke(Color.black, lineWidth: 2))
                    
                        .padding(.horizontal, 125)
                        .padding(.top, 30)
                }
        }
        
        .padding()
        
    }
    
    
}

