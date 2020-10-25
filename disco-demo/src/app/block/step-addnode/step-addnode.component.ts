import { Component, OnInit, Output,Inject,EventEmitter} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {BlockService} from '../blockservice';

@Component({
  selector: 'step-addnode',
  templateUrl: './step-addnode.component.html',
  styleUrls: ['./step-addnode.component.scss']
})
export class AddNodeStepComponent implements OnInit {
  public blockService:BlockService;
  constructor(blockService:BlockService) {
    this.blockService = blockService;
  }
  ngOnInit() {
  }

  private gridApi;
  private gridColumnApi;

  public defaultColDef =
  {
    initialWidth: 100,
    sortable: true,
    resizable: true,
  };
  public rowData:any[];

  onGridReady(params) {
    this.gridApi = params.api;
    this.gridColumnApi = params.columnApi;
    this.queryCompany();
  }


  loadDynamicColumns() {
    let columns =[
      {field:"汽车厂商",width:200},
      {field:"零部件",width:200},
      {field:"供货商",width:200},
      {field:"供货时间",width:200},
      {field:"供货数量",width:200}
    ]

    this.gridApi.setColumnDefs(columns);
  }

  queryCompany(){
     this.loadDynamicColumns();
     this.blockService.getUpstream()
           .subscribe(
               gRes =>{
                    this.rowData = gRes;
               }
         );
  }


}

