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
    
    //@Binding var groupName: String
    var groupName = "TEST" // for debug purposes
    var groupMembers = ["Eetu", "Chris", "Jan"] // for debug purposes
    
    @State var paidForMembers: [Bool]
    @State var selectedMember: String?
    
    init() {
            
        _paidForMembers = State(initialValue: Array(repeating: false, count: groupMembers.count))
        }
    
    
    var body: some View {
        
        
        NavigationView {
            VStack{
                List {
                    
                    Section(header: Text("Expense")) {
                        HStack {
                            Text("Item:")
                                .frame(width: 100, alignment: .leading)
                            TextField("Enter value", text: $item)
                                .textFieldStyle(RoundedBorderTextFieldStyle())
                        }
                        
                        HStack {
                            Text("Price:")
                                .frame(width: 100, alignment: .leading)
                            TextField("Enter value", text: $price)
                                .textFieldStyle(RoundedBorderTextFieldStyle())
                        }
                    }
                    
                    Section(header: Text("Paid by")) {
                        ForEach(groupMembers, id: \.self) { member in
                            HStack {
                                Text(member)
                                    .frame(width: 100, alignment: .leading)
                                
                                Spacer()
                                
                                if selectedMember == member {
                                    Image(systemName: "checkmark")
                                        .foregroundColor(.blue) // Color for checkmark
                                }
                            }
                            .contentShape(Rectangle())
                            .onTapGesture {
                                if selectedMember == member {
                                    selectedMember = nil
                                } else {
                                    selectedMember = member
                                }
                            }
                        }
                    }
                    
                    Section(header: Text("Paid for")) {
                        ForEach(0..<groupMembers.count, id: \.self) { index in
                            HStack {
                                Text(groupMembers[index])
                                    .frame(width: 100, alignment: .leading) // Adjust as needed
                                
                                Spacer()
                                
                                Toggle("", isOn: $paidForMembers[index])
                                    .labelsHidden()
                            }
                        }
                    }
                }
            }
        }
        .navigationTitle(Text("Adding to \(groupName)"))
        .onAppear {
            if paidForMembers.count != groupMembers.count {
                paidForMembers = Array(repeating: false, count: groupMembers.count)
            }
        }
        
//        VStack(alignment: .leading){
//            
//            Text("Add new item")
//            
//            TextField("Item", text: $item)
//                .padding(15) // Internal padding inside the text field
//                .background(Color("Primary colour")) // Background color for the text field
//                .cornerRadius(10) // Corner radius for rounded edges
//                .overlay( // Overlay the border using the same corner radius
//                    RoundedRectangle(cornerRadius: 10)
//                        .stroke(Color.black, lineWidth: 2) // Reference custom colour from assets
//                )
//                .padding(10)
//            
//            Text("Price")
//            
//            
//            TextField("€ £ $", text: $price)
//                .keyboardType(.decimalPad)
//                .padding(15)
//                .background(Color("Primary colour"))
//                .cornerRadius(10)
//                .overlay(
//                    RoundedRectangle(cornerRadius: 10)
//                        .stroke(Color.black, lineWidth: 2)
//                )
//                .padding(10)
//                .padding(.horizontal, 15)
//            
//            
//            DatePicker(
//                "Date",
//                selection: $date,
//                displayedComponents: [.date]
//            )
//            .padding(.vertical)
//            
//            
//            
//            HStack{
//                
//                Text("Tag")
//                
//                Picker("Tag", selection: $tagindex, content: {
//                    Text("Travel").tag(0)
//                    Text("Food").tag(1)
//                    Text("House").tag(2)
//                })
//                .padding(10)
//                .background(Color.white)
//                .cornerRadius(10)
//                .overlay(
//                    RoundedRectangle(cornerRadius: 10)
//                        .stroke(Color.black, lineWidth: 2))
//                .padding(.horizontal,80)
//
//                Group {
//                    if tagindex == 0 {
//                        Image(systemName: "airplane")
//                            .font(.system(size: 30))
//                    } else if tagindex == 1 {
//                        Image(systemName: "fork.knife")
//                            .font(.system(size: 30))
//                    } else if tagindex == 2 {
//                        Image(systemName: "house.fill")
//                            .font(.system(size: 30))
//                    }
//                }
//                Spacer()
//            }
//            // end of Hstack
//            
//            
//            Button(
//                action: {UIApplication.shared.sendAction(#selector(UIResponder.resignFirstResponder), to: nil,
//                                                         from: nil, for: nil)}) {
//                    Label("Done", systemImage: "chevron.right.circle")
//                        .padding()
//                        .background(Color("Primary colour"))
//                        .foregroundColor(.black)
//                        .cornerRadius(10)
//                        .overlay(RoundedRectangle(cornerRadius: 10)
//                        .stroke(Color.black, lineWidth: 2))
//                    
//                        .padding(.horizontal, 125)
//                        .padding(.top, 30)
//                }
//        }
//        
//        .padding()
        
    }
    
    
}

