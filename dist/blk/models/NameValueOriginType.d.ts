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
import type { NameValueModuleType } from './NameValueModuleType';
import type { WarningType } from './WarningType';
/**
 * Contains origin details.
 * @export
 * @interface NameValueOriginType
 */
export interface NameValueOriginType {
    /**
     *
     * @type {NameValueModuleType}
     * @memberof NameValueOriginType
     */
    originName?: NameValueModuleType;
    /**
     * Contains destination column for Origin.
     * @type {string}
     * @memberof NameValueOriginType
     */
    destination?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof NameValueOriginType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof NameValueOriginType
     */
    type?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof NameValueOriginType
     */
    idContext?: string;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof NameValueOriginType
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the NameValueOriginType interface.
 */
export declare function instanceOfNameValueOriginType(value: object): boolean;
export declare function NameValueOriginTypeFromJSON(json: any): NameValueOriginType;
export declare function NameValueOriginTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NameValueOriginType;
export declare function NameValueOriginTypeToJSON(value?: NameValueOriginType | null): any;
