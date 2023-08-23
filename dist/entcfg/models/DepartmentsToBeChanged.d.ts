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
import type { DepartmentType } from './DepartmentType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for updating departments.
 * @export
 * @interface DepartmentsToBeChanged
 */
export interface DepartmentsToBeChanged {
    /**
     * Collection of departments.
     * @type {Array<DepartmentType>}
     * @memberof DepartmentsToBeChanged
     */
    departments?: Array<DepartmentType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DepartmentsToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof DepartmentsToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the DepartmentsToBeChanged interface.
 */
export declare function instanceOfDepartmentsToBeChanged(value: object): boolean;
export declare function DepartmentsToBeChangedFromJSON(json: any): DepartmentsToBeChanged;
export declare function DepartmentsToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): DepartmentsToBeChanged;
export declare function DepartmentsToBeChangedToJSON(value?: DepartmentsToBeChanged | null): any;
