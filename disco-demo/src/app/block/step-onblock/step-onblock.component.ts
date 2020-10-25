import { Component, OnInit, Output,Inject,EventEmitter} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {BlockService} from '../blockservice';


@Component({
  selector: 'step-onblock',
  templateUrl: './step-onblock.component.html',
  styleUrls: ['./step-onblock.component.scss']
})
export class OnBlockStepComponent implements OnInit {
  public blockService:BlockService;
  constructor(blockService:BlockService) {
    this.blockService = blockService;
  }
  ngOnInit() {
  }

  private gridApiV2;
  private gridColumnApiV2;

  public defaultColDef =
  {
    initialWidth: 100,
    sortable: true,
    resizable: true,
  };

  public rowDataV2:any[];


  onGridReadyV2(params) {
    this.gridApiV2 = params.api;
    this.gridColumnApiV2 = params.columnApi;
    this.queryCompanyV2();
  }

  loadDynamicColumnsV2() {
      let columnsV2 =[
        {field:"汽车厂商",width:200},
        {field:"车型",width:200},
        {field:"经销商",width:200},
        {field:"交货时间",width:200},
        {field:"交货数量",width:200}
      ]

      this.gridApiV2.setColumnDefs(columnsV2);
  }

  queryCompanyV2(){
     this.loadDynamicColumnsV2();
     this.blockService.getDownstream()
           .subscribe(
               gRes =>{
                    this.rowDataV2 = gRes;
               }
         );
  }
  onBoarding(){
    window.open("www.baidu.com","_blank")
  }
}

