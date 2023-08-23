/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Information regarding why reservation is being /has been waitlisted.
 * @export
 * @interface WaitlistResType
 */
export interface WaitlistResType {
    /**
     * Waitlist Reason Description.
     * @type {string}
     * @memberof WaitlistResType
     */
    reasonDescription?: string;
    /**
     * Waitlist priority Description.
     * @type {string}
     * @memberof WaitlistResType
     */
    priorityDescription?: string;
    /**
     * Description why the reservation is being /has been waitlisted.
     * @type {string}
     * @memberof WaitlistResType
     */
    description?: string;
    /**
     * Waitlist Reason Code.
     * @type {string}
     * @memberof WaitlistResType
     */
    reasonCode?: string;
    /**
     * Waitlist Priority Code.
     * @type {string}
     * @memberof WaitlistResType
     */
    priorityCode?: string;
    /**
     * Guest telephone number.
     * @type {string}
     * @memberof WaitlistResType
     */
    telephone?: string;
}
/**
 * Check if a given object implements the WaitlistResType interface.
 */
export declare function instanceOfWaitlistResType(value: object): boolean;
export declare function WaitlistResTypeFromJSON(json: any): WaitlistResType;
export declare function WaitlistResTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): WaitlistResType;
export declare function WaitlistResTypeToJSON(value?: WaitlistResType | null): any;
