//student enters the name
//then writes the discipline and grade for it (A,A-,B,B-,C,C-,D,D-,F)
//then programm reads it and calculate the overall gpa
package main

import "fmt"

func main() {

	type Discipline struct{
		Credits int
		Grade string
		Points float64
	}

	gpa_map := map[string]float64{
		"A": 4.00,
		"A-": 3.50,
		"B": 3.00,
		"B-": 2.80,
		"C": 2.50,
		"C-": 2.00,
		"D": 1.85,
	}

	var name string
	var numberOfDisciplines int
	grades := map[string]Discipline{}


	fmt.Println("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil{
		fmt.Println("error", err)
		return
	}


	fmt.Println("Enter number of disciplines: ")
	_, err1 := fmt.Scan(&numberOfDisciplines)
	if err1 != nil{
		fmt.Println("error ", err)
		return
	}

	for i:=0; i<numberOfDisciplines; i++{
		var disc_name string
		var d Discipline
		fmt.Println("Enter discipline, number of credist and 1)grade letter 2)points")
		_, err := fmt.Scan(&disc_name, &d.Credits, &d.Grade, &d.Points)
		if err != nil{
			fmt.Println("Error ", err)
			return
		}
		grades[disc_name] = d
	}

	
	var sum_of_credits int
	var upper_sum float64
	for _, value := range grades{
		credits := value.Credits
		grade := value.Grade

		sum_of_credits += credits
		upper_sum += gpa_map[grade]* float64(credits)
	}

	final_gpa := upper_sum/float64(sum_of_credits)
	fmt.Println("Your GPA: ", final_gpa)

}
