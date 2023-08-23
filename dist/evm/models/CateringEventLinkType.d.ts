/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CateringEventLinkTypeType } from './CateringEventLinkTypeType';
import type { DateTimeSpanType } from './DateTimeSpanType';
import type { UniqueIDType } from './UniqueIDType';
/**
 *
 * @export
 * @interface CateringEventLinkType
 */
export interface CateringEventLinkType {
    /**
     *
     * @type {UniqueIDType}
     * @memberof CateringEventLinkType
     */
    id?: UniqueIDType;
    /**
     *
     * @type {CateringEventLinkTypeType}
     * @memberof CateringEventLinkType
     */
    type?: CateringEventLinkTypeType;
    /**
     *
     * @type {DateTimeSpanType}
     * @memberof CateringEventLinkType
     */
    masterEventDates?: DateTimeSpanType;
}
/**
 * Check if a given object implements the CateringEventLinkType interface.
 */
export declare function instanceOfCateringEventLinkType(value: object): boolean;
export declare function CateringEventLinkTypeFromJSON(json: any): CateringEventLinkType;
export declare function CateringEventLinkTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CateringEventLinkType;
export declare function CateringEventLinkTypeToJSON(value?: CateringEventLinkType | null): any;
