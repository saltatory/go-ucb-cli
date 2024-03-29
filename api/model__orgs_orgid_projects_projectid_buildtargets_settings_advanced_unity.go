/*
 * Unity Cloud Build
 *
 * This API is intended to be used in conjunction with the Unity Cloud Build service. A tool for building your Unity projects in the Cloud.  See https://developer.cloud.unity3d.com for more information.  ## Making requests This website is built to allow requests to be made against the API. If you are currently logged into Cloud Build you should be able to make requests without entering an API key.   You can find your API key in the Unity Cloud Services portal by clicking on 'Cloud Build Preferences' in the sidebar. Copy the API Key and paste it into the upper left corner of this website. It will be used in all subsequent requests.  ## Clients The Unity Cloud Build API is based upon Swagger. Client libraries to integrate with your projects can easily be generated with the [Swagger Code Generator](https://github.com/swagger-api/swagger-codegen).  The JSON schema required to generate a client for this API version is located here:  ``` [API_URL][BASE_PATH]/api.json ```  ## Authorization The Unity Cloud Build API requires an access token from your Unity Cloud Build account, which can be found at https://build.cloud.unity3d.com/login/me  To authenticate requests, include a Basic Authentication header with your API key as the value. e.g.  ``` Authorization: Basic [YOUR API KEY] ```  ## Pagination Paged results will take two parameters. A page number that is calculated based upon the per_page amount. For instance if there are 40 results and you specify page 2 with per_page set to 10 you will receive records 11-20.  Paged results will also return a Content-Range header. For the example above the content range header would look like this:  ``` Content-Range: items 11-20/40 ```  ## Versioning The API version is indicated in the request URL. Upgrading to a newer API version can be done by changing the path.  The API will receive a new version in the following cases:    * removal of a path or request type   * addition of a required field   * removal of a required field  The following changes are considered backwards compatible and will not trigger a new API version:    * addition of an endpoint or request type   * addition of an optional field   * removal of an optional field   * changes to the format of ids  ## Rate Limiting Requests against the Cloud Build API are limited to a rate of 100 per minute. To preserve the quality of service throughout Cloud Build, additional rate limits may apply to some actions. For example, polling aggressively instead of using webhooks or making API calls with a high concurrency may result in rate limiting.  It is not intended for these rate limits to interfere with any legitimate use of the API. Please contact support at <cloudbuild@unity3d.com> if your use is affected by this rate limit.  You can check the returned HTTP headers for any API request to see your current rate limit status.   * __X-RateLimit-Limit:__ maximum number of requests per minute   * __X-RateLimit-Remaining:__ remaining number of requests in the current window   * __X-RateLimit-Reset:__ time at which the current window will reset (UTC epoch seconds)  Once you go over the rate limit you will receive an error response: ``` HTTP Status: 429 {   \"error\": \"Rate limit exceeded, retry in XX seconds\" } ``` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ucbapi
// OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnity struct for OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnity
type OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnity struct {
	// The fully-qualified name of a public static method you want us to call before we start the Unity build process. For example: ClassName.NeatMethod or NameSpace.ClassName.NeatMethod. No trailing parenthesis, and it can't have the same name as your Post-Export method!
	PreExportMethod string `json:"preExportMethod,omitempty"`
	// The fully-qualified name of a public static method you want us to call after we finish the Unity build process (but before Xcode). For example: ClassName.CoolMethod or NameSpace.ClassName.CoolMethod. No trailing parenthesis, and it can't have the same name as your Post-Export method! This method must accept a string parameter, which will receive the path to the exported Unity player (or Xcode project in the case of iOS).
	PostExportMethod string `json:"postExportMethod,omitempty"`
	// Relative path to the script that should be run before the build process starts.
	PreBuildScript string `json:"preBuildScript,omitempty"`
	// Relative path to the script that should be run after the build process finishes.
	PostBuildScript string `json:"postBuildScript,omitempty"`
	// Enter the names of the symbols you want to define for iOS. These symbols can then be used as the conditions for #if directives just like the built-in ones. (i.e. #IF MYDEFINE or #IF AMAZON)
	ScriptingDefineSymbols string `json:"scriptingDefineSymbols,omitempty"`
	PlayerExporter OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityPlayerExporter `json:"playerExporter,omitempty"`
	PlayerSettings OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityPlayerSettings `json:"playerSettings,omitempty"`
	EditorUserBuildSettings OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityEditorUserBuildSettings `json:"editorUserBuildSettings,omitempty"`
	AssetBundles OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityAssetBundles `json:"assetBundles,omitempty"`
	Addressables OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityAddressables `json:"addressables,omitempty"`
	// Run any unit tests your project has when a build happens.
	RunUnitTests bool `json:"runUnitTests,omitempty"`
	// Should Edit Mode unit tests be run? NOTE: requires runUnitTests to be true and building with Unity 5.6 or newer.
	RunEditModeTests bool `json:"runEditModeTests,omitempty"`
	// Should Play Mode unit tests be run? NOTE: requires runUnitTests to be true and building with Unity 5.6 or newer.
	RunPlayModeTests bool `json:"runPlayModeTests,omitempty"`
	// Mark builds as failed if the unit tests do not pass.
	FailedUnitTestFailsBuild bool `json:"failedUnitTestFailsBuild,omitempty"`
	// LEGACY - The Unity method to call when running unit tests (only supported in Unity 5.2 and lower).
	UnitTestMethod string `json:"unitTestMethod,omitempty"`
	// Enable lightmap baking (disabled by default since it is very slow and usually unnecessary)
	EnableLightBake bool `json:"enableLightBake,omitempty"`
}
