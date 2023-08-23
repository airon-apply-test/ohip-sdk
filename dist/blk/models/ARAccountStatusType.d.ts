/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Type for the Account Status. Status can be restricted and a message/description added to the Account.
 * @export
 * @interface ARAccountStatusType
 */
export interface ARAccountStatusType {
    /**
     * The Restriction Code added on the Account. This is available when the functionality for adding restriction codes is ON.
     * @type {string}
     * @memberof ARAccountStatusType
     */
    restriction?: string;
    /**
     * User defined status message on the Account.
     * @type {string}
     * @memberof ARAccountStatusType
     */
    description?: string;
    /**
     * Flag to indicate if the Account is restricted.
     * @type {boolean}
     * @memberof ARAccountStatusType
     */
    restricted?: boolean;
}
/**
 * Check if a given object implements the ARAccountStatusType interface.
 */
export declare function instanceOfARAccountStatusType(value: object): boolean;
export declare function ARAccountStatusTypeFromJSON(json: any): ARAccountStatusType;
export declare function ARAccountStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARAccountStatusType;
export declare function ARAccountStatusTypeToJSON(value?: ARAccountStatusType | null): any;
