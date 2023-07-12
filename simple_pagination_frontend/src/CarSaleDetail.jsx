import TableView from "./TableView"

export default function () {
	const saleData = [
		{"id":1,"full_name":"Salomi Glyne","email":"sglyne0@europa.eu","gender":"Female","car":{"make":"Toyota","model":"Tacoma Xtra","vin":"WBAKB8C56AC955228"}},
		{"id":2,"full_name":"Kayne Linthead","email":"klinthead1@vistaprint.com","gender":"Male","car":{"make":"Mercedes-Benz","model":"E-Class","vin":"WAUFFAFL6BA647951"}}
	]

	return (
		<>
			<TableView tableData={saleData} />
		</>
	)
}
