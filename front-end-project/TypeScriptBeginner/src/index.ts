// let age=20; //int Type
// let num=123_456_789; //int Type
// let Name='xyz'; //string Type
// let id; //any Type variables
//  id=89;
//  id='ryj'

//function
//  function render(document: any){
//     console.log(document);
//  }
//arrays

//let arr=[1,2,3,4]; //in js

//let arr1: number[]=[1,2,3];//in ts

// let arr2:number[]=[]
// arr2[0]=1;
// arr2[1]='2'

/*let x: number[]=[];      //x is a variable and number[] is a type of integer type array
x.forEach(n=>n.toString)*/ //n=>n arrow function

//tuples ..basicall tuples works on fixed length array

/*let user:[number,string]=[1,'Physics'];*/

//Pascal Case
//enums

/*enum Size { Small=1, Medium, Large } //will give lengthy js code
let mySize: Size = Size.Medium
console.log(mySize);*/

/*const enum Size { Small=1, Medium, Large } //will give optimize js code only by using const keyword
let mySize: Size = Size.Medium
console.log(mySize);*/

//functions

// function avg(a:number):number{
//    return a;

// }

//objects

// let employee:{
//    id:number,
//    name:string,
// }={
//    id:1,
//    name:'hnj',
// }

//Advanced Types

// type Employee={
//  id:number,
//  name:string,
//  retire: (date:Date)=> void
// }

// let employee:Employee={
//  id:1,
//  name:'tjnt',
//  retire:(date:Date)=>{
//    console.log(date);
//  }
// }


//unions

function kgToLbs(weight: number | string): number {

   //  weight.toLocaleString,
   //  weight.toString,
   //  weight.valueOf,

   //Narrowing

   if (typeof weight === 'number') {
      //if we declare that weight is type of number all 
      //the method related to numbers are available

      // weight.toExponential
      // weight.toFixed
      // weight.toLocaleString
      // weight.toPrecision
      // weight.toString
      // weight.valueOf

      return weight * 2.2;

   }
   else {
      return parseInt(weight) * 2.2 ;
   }
}
kgToLbs(10);
kgToLbs('10kg');

