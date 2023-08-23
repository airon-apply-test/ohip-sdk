/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AutomaticTransmissionScheduleType } from './AutomaticTransmissionScheduleType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PostAutomaticTransmissionSchedulesRequest
 */
export interface PostAutomaticTransmissionSchedulesRequest {
    /**
     * List of automatic transmission schedules
     * @type {Array<AutomaticTransmissionScheduleType>}
     * @memberof PostAutomaticTransmissionSchedulesRequest
     */
    automaticTransmissionSchedules?: Array<AutomaticTransmissionScheduleType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostAutomaticTransmissionSchedulesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostAutomaticTransmissionSchedulesRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostAutomaticTransmissionSchedulesRequest interface.
 */
export declare function instanceOfPostAutomaticTransmissionSchedulesRequest(value: object): boolean;
export declare function PostAutomaticTransmissionSchedulesRequestFromJSON(json: any): PostAutomaticTransmissionSchedulesRequest;
export declare function PostAutomaticTransmissionSchedulesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostAutomaticTransmissionSchedulesRequest;
export declare function PostAutomaticTransmissionSchedulesRequestToJSON(value?: PostAutomaticTransmissionSchedulesRequest | null): any;
