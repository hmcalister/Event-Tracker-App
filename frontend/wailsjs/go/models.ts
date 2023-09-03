export namespace models {
	
	export class Event {
	    Name: string;
	    StartTime: Date;
	    TimeoutDuration: number;
	    IsRecurring: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Event(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.StartTime = new Date(source["StartTime"]);
	        this.TimeoutDuration = source["TimeoutDuration"];
	        this.IsRecurring = source["IsRecurring"];
	    }
	}

}

