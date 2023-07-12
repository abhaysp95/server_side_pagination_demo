import { Table } from "react-bootstrap";

export default function({tableData}) {
	return <>
			<Table striped bordered hover>
				<thead>
					<tr>
					{
						Object.keys(tableData[0]).map((saleKey, idx) => (
							<th key={idx} className="text-center">{saleKey}</th>
						))
					}
					</tr>
				</thead>
				<tbody>
				{
					tableData.map((saleRow, idx) => (
						<tr key={`sale_row_${idx}`}>
						{
							Object.entries(saleRow).map(([saleKey, saleCol], cidx) => {
								let colData = saleKey === "car"
									? `${saleCol.make} ${saleCol.model} [${saleCol.vin}]`
									: saleCol
								return (
									<td
										className="text-center"
										key={`sale_col_${idx}_${cidx}`}
									>{colData}</td>
								)
							})
						}
						</tr>
					))
				}
				</tbody>
			</Table>
		</>
}
