export namespace main {
	
	export class NumberObj {
	    NewName: string;
	    Suffix: number;
	
	    static createFrom(source: any = {}) {
	        return new NumberObj(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.NewName = source["NewName"];
	        this.Suffix = source["Suffix"];
	    }
	}
	export class RenamePreview {
	    OldDisPlayName: string;
	    NewDisPlayName: string;
	    OldName: string;
	    NewName: string;
	
	    static createFrom(source: any = {}) {
	        return new RenamePreview(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.OldDisPlayName = source["OldDisPlayName"];
	        this.NewDisPlayName = source["NewDisPlayName"];
	        this.OldName = source["OldName"];
	        this.NewName = source["NewName"];
	    }
	}
	export class ReplaceObj {
	    OldStr: string;
	    NewStr: string;
	
	    static createFrom(source: any = {}) {
	        return new ReplaceObj(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.OldStr = source["OldStr"];
	        this.NewStr = source["NewStr"];
	    }
	}
	export class Rules {
	    RenameType: number;
	    Prefix: string;
	    Suffix: string;
	    ReplaceObj?: ReplaceObj;
	    NumberObj?: NumberObj;
	
	    static createFrom(source: any = {}) {
	        return new Rules(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RenameType = source["RenameType"];
	        this.Prefix = source["Prefix"];
	        this.Suffix = source["Suffix"];
	        this.ReplaceObj = this.convertValues(source["ReplaceObj"], ReplaceObj);
	        this.NumberObj = this.convertValues(source["NumberObj"], NumberObj);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

