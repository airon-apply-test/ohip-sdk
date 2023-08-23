/**
 * OPERA Cloud Leisure Management API
 * APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ActivityBookingRQType } from './ActivityBookingRQType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PutActivityBookingRequest
 */
export interface PutActivityBookingRequest {
    /**
     *
     * @type {ActivityBookingRQType}
     * @memberof PutActivityBookingRequest
     */
    activityBooking?: ActivityBookingRQType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PutActivityBookingRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutActivityBookingRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PutActivityBookingRequest interface.
 */
export declare function instanceOfPutActivityBookingRequest(value: object): boolean;
export declare function PutActivityBookingRequestFromJSON(json: any): PutActivityBookingRequest;
export declare function PutActivityBookingRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutActivityBookingRequest;
export declare function PutActivityBookingRequestToJSON(value?: PutActivityBookingRequest | null): any;
