/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ActivityDetailsType } from './ActivityDetailsType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface Activity
 */
export interface Activity {
    /**
     *
     * @type {ActivityDetailsType}
     * @memberof Activity
     */
    activityDetails?: ActivityDetailsType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof Activity
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof Activity
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the Activity interface.
 */
export declare function instanceOfActivity(value: object): boolean;
export declare function ActivityFromJSON(json: any): Activity;
export declare function ActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): Activity;
export declare function ActivityToJSON(value?: Activity | null): any;
