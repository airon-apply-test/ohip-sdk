/**
 * OPERA Cloud Price Availability Rate Async API
 * APIs to cater for Price and Rate Availability Asynchronous functionality in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RestrictionStatusesType } from './RestrictionStatusesType';
/**
 * The RestrictionStatusType is used to indicate the type of restriction applied. An enumerated restriction type is defined in the attribute group.
 * @export
 * @interface RestrictionStatusType
 */
export interface RestrictionStatusType {
    /**
     *
     * @type {RestrictionStatusesType}
     * @memberof RestrictionStatusType
     */
    code?: RestrictionStatusesType;
    /**
     * Defines restriction in conjunction with Restriction Code. Value must be specified for Restriction Codes MinimumStayThrough, MaximumStayThrough, MinimumLengthOfStay, MaximumLengthOfStay, MinimumAdvanceBooking, MaximumAdvanceBooking.
     * @type {number}
     * @memberof RestrictionStatusType
     */
    unit?: number;
    /**
     * Indicates Length of Stay 1 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay1?: boolean;
    /**
     * Indicates Length of Stay 2 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay2?: boolean;
    /**
     * Indicates Length of Stay 3 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay3?: boolean;
    /**
     * Indicates Length of Stay 4 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay4?: boolean;
    /**
     * Indicates Length of Stay 5 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay5?: boolean;
    /**
     * Indicates Length of Stay 6 is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay6?: boolean;
    /**
     * Indicates Length of Stay 7 or more is not available. Defines restriction in conjunction with Restriction Code LengthOfStayNotAvailable.
     * @type {boolean}
     * @memberof RestrictionStatusType
     */
    lengthOfStay7?: boolean;
}
/**
 * Check if a given object implements the RestrictionStatusType interface.
 */
export declare function instanceOfRestrictionStatusType(value: object): boolean;
export declare function RestrictionStatusTypeFromJSON(json: any): RestrictionStatusType;
export declare function RestrictionStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RestrictionStatusType;
export declare function RestrictionStatusTypeToJSON(value?: RestrictionStatusType | null): any;
