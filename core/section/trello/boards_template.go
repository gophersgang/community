// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package trello

const boardsTemplate = `
<div class="section-trello-render">
	{{if gt (len .Boards) 0}}
		<div class="heading">Boards</div>
		<p>Changes since {{.Since}}.</p>
		<div class="section-trello-render">
			<table class="trello-table" class="width-100">
				<tbody class="board-stats">
					<tr>
						<td>
							<span class="stat-number">{{len .Boards}}</span>boards
						</td>
						<td>
							<span class="stat-number">{{.ListTotal}}</span>lists
						</td>
						<td>
							<span class="stat-number">{{.CardTotal}}</span>cards
						</td>
						<td>
							<span class="stat-number">{{len .MemberBoardAssign}}</span>members
						</td>
					</tr>
				</tbody>
			</table>
			<table class="trello-table" class="width-100">
				<tbody class="trello">
				{{range $b := .Boards}}
					<tr>
						<td>
							<a href="{{ $b.Board.URL }}">
								<span class="trello-board" style="background-color: {{$b.Board.Prefs.BackgroundColor}}">{{$b.Board.Name}}</span>
							</a>
						</td>
						<td>
							<div class="board-summary">There are {{ len $b.Actions }} actions for this board.</div>
							<span class="board-meta">
								{{range $act, $tot := $b.ActionSummary}}
									{{$act}} ({{$tot}}),
								{{end}}
								{{if gt (len $b.Archived) 0}}
									archive card ({{len $b.Archived}}).
								{{else}}
									no cards archived.
								{{end}}
								<br>
							</span>
						</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	{{end}}
</div>
`