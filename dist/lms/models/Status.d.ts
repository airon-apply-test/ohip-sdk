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
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Existing Operations Responses will eventually be modified to be extended from this type.
 * @export
 * @interface Status
 */
export interface Status {
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof Status
     */
    warnings?: Array<WarningType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof Status
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the Status interface.
 */
export declare function instanceOfStatus(value: object): boolean;
export declare function StatusFromJSON(json: any): Status;
export declare function StatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): Status;
export declare function StatusToJSON(value?: Status | null): any;
