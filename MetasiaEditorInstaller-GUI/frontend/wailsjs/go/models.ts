export namespace main {
	
	export class PluginInfo {
	    AssetUrl: string;
	    FileName: string;
	
	    static createFrom(source: any = {}) {
	        return new PluginInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.AssetUrl = source["AssetUrl"];
	        this.FileName = source["FileName"];
	    }
	}
	export class InstallParams {
	    Path: string;
	    MetasiaAssetsUrl: string;
	    MetasiaAssetFileName: string;
	    PluginsPath: string;
	    Plugins: PluginInfo[];
	
	    static createFrom(source: any = {}) {
	        return new InstallParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.MetasiaAssetsUrl = source["MetasiaAssetsUrl"];
	        this.MetasiaAssetFileName = source["MetasiaAssetFileName"];
	        this.PluginsPath = source["PluginsPath"];
	        this.Plugins = this.convertValues(source["Plugins"], PluginInfo);
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

