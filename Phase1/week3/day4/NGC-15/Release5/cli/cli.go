package cli

import (
	"Phase1/week3/day4/NGC-15/Release5/config"
	"Phase1/week3/day4/NGC-15/Release5/entity"
	"Phase1/week3/day4/NGC-15/Release5/handler"
	"bufio"
	"fmt"
	"os"
	"time"
)

func RunProgram() {

	db, err := config.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	var input string
	var firstname, lastname, position string
	var menu_name, description, category string
	var price float64
	var table_number, employee_id int
	var order_date, status string
	now := time.Now()
	var item_id, qty int
	var subtotal float64
	var paymentmethod, paymentdate, paymentconf string

	var tempTableNum int
	var contInputOrder int

	for {
		fmt.Println("============\tWelcome to Tasty Bites\t============")
		fmt.Println("============ \tEmployee Management\t============")
		fmt.Println("1. Add new employee")
		fmt.Println("2. View all list of employees")

		fmt.Println("============ Menu Management ============")
		fmt.Println("3. Add new menu items")
		fmt.Println("4. View a list of all menu items")

		fmt.Println("============ Order Processing ============")
		fmt.Println("5. Create a new order")
		fmt.Println("6. Add menu items to an order and calculate the price")
		fmt.Println("7. View a list of all orders")

		fmt.Println("============ Payment Processing ============")
		fmt.Println("8. Process and record payment for an order")
		fmt.Println("9. Exit")

		fmt.Print("Input > ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("ERROR INPUT MENU : ", err.Error())
			return
		}

		switch input {
			case "1":
				fmt.Println("Input new employee data:")
				fmt.Print("Firstname > ")
				_, err := fmt.Fscan(reader, &firstname)
				if err != nil {
					reader.ReadString('\n')
				}

				fmt.Print("Lastname > ")
				_, err = fmt.Fscan(reader, &lastname)
				if err != nil {
					reader.ReadString('\n')
				}

				fmt.Print("Position > ")
				_, err = fmt.Fscan(reader, &position)
				if err != nil {
					reader.ReadString('\n')
				}

				var employee = entity.Employee{
					First_name: firstname,
					Last_name: lastname,
					Position: position,
				}
				err = handler.InsertNewEmployee(db, employee)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Printf("%s now become an employee\n\n", firstname)
			case "2":
				listEmployees, err := handler.GetAllEmployees(db)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println("List of all employees:")
				for i, v := range listEmployees {
					fmt.Printf("%d. %s %s\t : %s\n", i+1, v.First_name, v.Last_name, v.Position)
				}
			case "3":
				fmt.Println("Input new menu:")
				fmt.Print("Menu name > ")
				_, err := fmt.Fscan(reader, &menu_name)
				if err != nil {
					reader.ReadString('\n')
					fmt.Println(err.Error())
				}
				
				fmt.Print("Description > ")
				_, err = fmt.Fscan(reader, &description)
				if err != nil {
					reader.ReadString('\n')
					fmt.Println(err.Error())
				}

				fmt.Print("Price > ")
				_, err = fmt.Scanf("%f", &price)
				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Print("Category > ")
				_, err = fmt.Fscan(reader, &category)
				if err != nil {
					reader.ReadString('\n')
					fmt.Println(err.Error())
				}

				var menu = entity.MenuItem {
					Name: menu_name,
					Description: description,
					Price: price,
					Category: category,
				}
				err = handler.InsertNewMenu(db, menu)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Printf("%s now become a menu\n\n", menu_name)
			case "4":
				listMenuItems, err := handler.GetAllMenuItems(db)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println("List of all menus:")
				for i, v := range listMenuItems {
					fmt.Printf("%d. %s (%s)\n", i+1, v.Name, v.Category)
					fmt.Println("  ", v.Description)
					fmt.Printf("   Price: $%.2f\n", v.Price)
				}
			case "5":
				fmt.Println("Input new order:")
				fmt.Print("Table number > ")
				_, err = fmt.Scanln(&table_number)
				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Print("Employee id > ")
				_, err = fmt.Scanln(&employee_id)
				if err != nil {
					fmt.Println(err.Error())
				}

				order_date = now.Format("2006-01-02")

				fmt.Print("Status > ")
				_, err = fmt.Fscan(reader, &status)
				if err != nil {
					reader.ReadString('\n')
					fmt.Println(err.Error())
				}

				var order = entity.Order{
					Table_number: table_number,
					Employee_id: employee_id,
					Order_date: order_date,
					Status: status,
				}
				
				err = handler.CreateNewOrder(db, order)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println("Success to add the order")
			case "6":
				fmt.Print("Input your table number > ")
				_, err = fmt.Scanln(&tempTableNum)
				if err != nil {
					fmt.Println(err.Error())
				}

				orderID, err := handler.GetCurrentOrderID(db, tempTableNum)
				if err != nil || orderID == -1 {
					fmt.Println(err.Error())
					return 
				}
				for {
					fmt.Println("=== Input menu items to order: ===")
					fmt.Print("Menu ID > ")
					_, err = fmt.Scanln(&item_id)
					if err != nil {
						fmt.Println(err.Error())
					}

					fmt.Print("Quantity > ")
					_, err = fmt.Scanln(&qty)
					if err != nil {
						fmt.Println(err.Error())
					}

					price, err := handler.GetItemPrice(db, item_id)
					if err != nil || price == -1 {
						fmt.Println(err.Error())
						return 
					}

					subtotal = price * float64(qty)

					var orderitem = entity.OrderItems{
						Order_id: orderID,
						Item_id: item_id,
						Quantity: qty,
						Subtotal: subtotal,
					}

					err = handler.CreateNewOrderItem(db, orderitem)
					if err != nil {
						fmt.Println(err.Error())
						return 
					}

					fmt.Println("Success to input your order")

					fmt.Print("Do you want to continue order (1/0) ? [Press 0 to exit] > ")
					_, err = fmt.Scanln(&contInputOrder)
					if err != nil {
						fmt.Println(err.Error())
					}

					if contInputOrder == 1 {
						continue
					}else{
						break
					}
				}
			case "7":
				orderdetails, err := handler.GetAllOrdersDetail(db)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println("List of all orders:")
				for _, v := range orderdetails {
					fmt.Printf("=== For table number %d ===\n", v.TableNumber)
					fmt.Printf("Order ID %d.  %s  $%.2f \n", v.OrderID, v.Name, v.TotalPrice)
					fmt.Println()
				}
			case "8":
				fmt.Print("Input your table number > ")
				_, err = fmt.Scanln(&tempTableNum)
				if err != nil {
					fmt.Println(err.Error())
				}

				orderID, err := handler.GetCurrentOrderID(db, tempTableNum)
				if err != nil || orderID == -1 {
					fmt.Println(err.Error())
					return 
				}

				fmt.Print("Payment Method > ")
				_, err = fmt.Fscan(reader, &paymentmethod)
				if err != nil {
					reader.ReadString('\n')
				}

				paymentdate = now.Format("2006-01-02")

				totalamount, err := handler.GetTotalAmountFromOrderID(db, orderID)
				if err != nil {
					fmt.Println(err.Error())
					return 
				}

				fmt.Print("You want to pay? (yes/no) > ")
				_, err = fmt.Fscan(reader, &paymentconf)
				if err != nil {
					reader.ReadString('\n')
				}

				if paymentconf == "yes" {
					var payment = entity.Payment{
						OrderID: orderID,
						PaymentDate: paymentdate,
						PaymentMethod: paymentmethod,
						TotalAmount: totalamount,
					}
					err = handler.InsertNewPaymentOrder(db, payment)
					if err != nil {
						fmt.Println(err.Error())
						return 
					}
					err = handler.ChangeStatusOrder(db, orderID)
					if err != nil {
						fmt.Println(err.Error())
						return 
					}
				}
			case "9":
				os.Exit(1)
		}
	}

}