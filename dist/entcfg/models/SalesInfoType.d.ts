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
 * Provides sales information about the profiles of type company, travel agent, source and contact.
 * @export
 * @interface SalesInfoType
 */
export interface SalesInfoType {
    /**
     * Defines the scope.
     * @type {string}
     * @memberof SalesInfoType
     */
    scope?: string;
    /**
     * Defines the scope city.
     * @type {string}
     * @memberof SalesInfoType
     */
    scopeCity?: string;
    /**
     * Defines the account type.
     * @type {string}
     * @memberof SalesInfoType
     */
    accountType?: string;
    /**
     * Defines the account source.
     * @type {string}
     * @memberof SalesInfoType
     */
    accountSource?: string;
    /**
     * Defines the industry code.
     * @type {string}
     * @memberof SalesInfoType
     */
    industryCode?: string;
    /**
     * Defines the Business segments.
     * @type {string}
     * @memberof SalesInfoType
     */
    businessSegments?: string;
    /**
     * Defines the priority.
     * @type {string}
     * @memberof SalesInfoType
     */
    priority?: string;
    /**
     * Defines the rooms potential.
     * @type {string}
     * @memberof SalesInfoType
     */
    roomsPotential?: string;
    /**
     * Defines the action code.
     * @type {string}
     * @memberof SalesInfoType
     */
    actionCode?: string;
    /**
     * Defines the competition code.
     * @type {string}
     * @memberof SalesInfoType
     */
    competitionCode?: string;
    /**
     * Defines the influence for the contact profile.
     * @type {string}
     * @memberof SalesInfoType
     */
    influence?: string;
    /**
     * Defines the Preferred Room for profile.
     * @type {string}
     * @memberof SalesInfoType
     */
    preferredRoom?: string;
    /**
     * Hotel Code used to filter the sales information.
     * @type {string}
     * @memberof SalesInfoType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the SalesInfoType interface.
 */
export declare function instanceOfSalesInfoType(value: object): boolean;
export declare function SalesInfoTypeFromJSON(json: any): SalesInfoType;
export declare function SalesInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SalesInfoType;
export declare function SalesInfoTypeToJSON(value?: SalesInfoType | null): any;
